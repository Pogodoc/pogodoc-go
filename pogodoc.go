package pogodoc

import (
	"context"
	"fmt"
	"os"

	"github.com/Pogodoc/pogodoc-go/client/client"
	"github.com/Pogodoc/pogodoc-go/client/option"
)

func PogodocClientInit() (*PogodocClient, error) {
	var tokenString string
	var baseURL string
	if os.Getenv("POGODOC_BASE_URL") != "" {
		baseURL = os.Getenv("POGODOC_BASE_URL")
	} else {
		baseURL = Environments.Default
	}
	if os.Getenv("POGODOC_API_TOKEN") != "" {
		tokenString = os.Getenv("POGODOC_API_TOKEN")

	} else {
		return nil, fmt.Errorf("API token is required. Please provide it either as a parameter or set the POGODOC_API_TOKEN environment variable")
	}
	c := client.NewClient(
		option.WithToken(tokenString),
		option.WithBaseURL(baseURL),
	)
	return &PogodocClient{Client: c}, nil
}

func PogodocClientInitWithConfig(baseURL string, tokenString string) (*PogodocClient, error) {

	c := client.NewClient(
		option.WithToken(tokenString),
		option.WithBaseURL(baseURL),
	)

	return &PogodocClient{Client: c}, nil
}

func PogodocClientInitWithToken(tokenString string) (*PogodocClient, error) {
	c := client.NewClient(
		option.WithToken(tokenString),
	)
	return &PogodocClient{Client: c}, nil
}

func (c *PogodocClient) SaveTemplate(filePath string, metadata SaveCreatedTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	payload, err := ReadFile(filePath)
	if err != nil {
		return "", err
	}
	payloadLength := len(payload)
	if payloadLength == 0 {
		return "", fmt.Errorf("error: File is empty")
	}

	fsProps := FileStreamProps{
		payload:       payload,
		payloadLength: payloadLength,
	}

	return c.SaveTemplateFromFileStream(fsProps, metadata, ctx)
}

func (c *PogodocClient) SaveTemplateFromFileStream(fsProps FileStreamProps, metadata SaveCreatedTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	response, err := c.Templates.InitializeTemplateCreation(ctx)
	if err != nil {
		return "", fmt.Errorf("initializing template creation: %v", err)
	}
	templateId := response.TemplateId

	err = UploadToS3WithURL(response.PresignedTemplateUploadUrl, fsProps, "application/zip")
	if err != nil {
		return "", fmt.Errorf("uploading template: %v", err)
	}

	err = c.Templates.ExtractTemplateFiles(ctx, templateId)
	if err != nil {
		return "", fmt.Errorf("extracting template files: %v", err)
	}
	request := GenerateTemplatePreviewsRequest{
		Type: GenerateTemplatePreviewsRequestType(metadata.Type),
		Data: metadata.SampleData,
	}

	previewResponse, err := c.Templates.GenerateTemplatePreviews(ctx, templateId, &request)
	if err != nil {
		return "", fmt.Errorf("generating template previews: %v", err)

	}
	previewPng := previewResponse.PngPreview.JobId
	previewPdf := previewResponse.PdfPreview.JobId

	saveCreatedTemplateRequest := SaveCreatedTemplateRequest{
		TemplateInfo: &SaveCreatedTemplateRequestTemplateInfo{
			Title:       metadata.Title,
			Description: metadata.Description,
			Categories:  metadata.Categories,
			Type:        metadata.Type,
			SourceCode:  metadata.SourceCode,
			SampleData:  metadata.SampleData,
		},
		PreviewIds: &SaveCreatedTemplateRequestPreviewIds{
			PngJobId: previewPng,
			PdfJobId: previewPdf,
		},
	}

	err = c.Templates.SaveCreatedTemplate(ctx, templateId, &saveCreatedTemplateRequest)
	if err != nil {
		return "", fmt.Errorf("saving created template: %v", err)
	}

	return templateId, nil

}

func (c *PogodocClient) UpdateTemplate(templateId string, filePath string, metadata UpdateTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	payload, err := ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("file is empty: %v", err)
	}

	payloadLength := len(payload)
	fsProps := FileStreamProps{
		payload:       payload,
		payloadLength: payloadLength,
	}

	return c.UpdateTemplateFromFileStream(templateId, fsProps, metadata, ctx)

}

func (c *PogodocClient) UpdateTemplateFromFileStream(templateId string, fsProps FileStreamProps, metadata UpdateTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	response, err := c.Templates.InitializeTemplateCreation(ctx)
	if err != nil {
		return "", fmt.Errorf("initializing template creation: %v", err)
	}
	contentId := response.TemplateId

	err = UploadToS3WithURL(response.PresignedTemplateUploadUrl, fsProps, "application/zip")
	if err != nil {
		return "", fmt.Errorf("uploading template: %v", err)
	}

	err = c.Templates.ExtractTemplateFiles(ctx, contentId)
	if err != nil {
		return "", fmt.Errorf("extracting template files: %v", err)
	}

	request := GenerateTemplatePreviewsRequest{
		Type: GenerateTemplatePreviewsRequestType(metadata.Type),
		Data: metadata.SampleData,
	}
	previewResponse, err := c.Templates.GenerateTemplatePreviews(ctx, contentId, &request)
	if err != nil {
		return "", fmt.Errorf("generating template previews: %v", err)
	}

	updateTemplateReq := &UpdateTemplateRequest{
		TemplateInfo: &UpdateTemplateRequestTemplateInfo{
			Title:       metadata.Title,
			Type:        metadata.Type,
			Description: metadata.Description,
			Categories:  metadata.Categories,
			SourceCode:  metadata.SourceCode,
			SampleData:  metadata.SampleData,
		},
		PreviewIds: &UpdateTemplateRequestPreviewIds{
			PngJobId: previewResponse.PngPreview.JobId,
			PdfJobId: previewResponse.PdfPreview.JobId,
		},
		ContentId: contentId,
	}

	_, err = c.Templates.UpdateTemplate(ctx, templateId, updateTemplateReq)
	if err != nil {
		return "", fmt.Errorf("updating template: %v", err)
	}

	return templateId, nil

}

func (c *PogodocClient) GenerateDocument(gdProps GenerateDocumentProps, ctx context.Context) (*GetJobStatusResponse, error) {

	initRequest := gdProps.InitializeRenderJobRequest
	initResponse, err := c.Documents.InitializeRenderJob(ctx, &initRequest)
	if err != nil {
		return nil, fmt.Errorf("initializing document render: %v", err)
	}

	Data := []byte(fmt.Sprint(gdProps.InitializeRenderJobRequest.Data))
	if initResponse != nil && initResponse.PresignedDataUploadUrl != nil {
		err = UploadToS3WithURL(*initResponse.PresignedDataUploadUrl, FileStreamProps{
			payload:       Data,
			payloadLength: len(Data),
		}, "application/json")
		if err != nil {
		}
	}

	template := gdProps.template
	if template != nil && *template != "" && initResponse.PresignedTemplateUploadUrl != nil {
		err = UploadToS3WithURL(*initResponse.PresignedTemplateUploadUrl, FileStreamProps{
			payload:       []byte(*template),
			payloadLength: len(*template),
		}, "text/html")
		if err != nil {
			return nil, fmt.Errorf("uploading document: %v", err)

		}
	}

	_, err = c.Documents.StartRenderJob(
		ctx,
		initResponse.JobId,
		&gdProps.StartRenderJobRequest,
	)
	if err != nil {
		return nil, fmt.Errorf("starting render: %v", err)
	}

	result, err := c.Documents.GetJobStatus(ctx, initResponse.JobId)
	if err != nil {
		return nil, fmt.Errorf("getting job status: %v", err)

	}

	return result, nil

}
