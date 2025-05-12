package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	api "pogodoc/go/sdk"
	"pogodoc/go/sdk/client"

	"github.com/joho/godotenv"
)

type PogodocEnv struct {
	baseURL string `env:"POGODOC_BASE_URL"`
	token   string `env:"POGODOC_TOKEN"`
}

func PogodocClientInit(baseURL string, tokenString string) PogodocClient {

	c := client.NewClient(
		client.ClientWithBaseURL(baseURL),
		client.ClientWithAuthToken(tokenString),
	)

	return PogodocClient{c}
}

func (c *PogodocClient) SaveTemplate(filePath string, metadata api.SaveCreatedTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	payload, err := ReadFile(filePath)
	if err != nil {
		return "", err
	}
	payloadLength := len(payload)
	if payloadLength == 0 {
		fmt.Println("Error: File is empty")
		return "", fmt.Errorf("error: File is empty")
	}

	fsProps := FileStreamProps{
		payload:       payload,
		payloadLength: payloadLength,
	}

	return c.SaveTemplateFromFileStream(fsProps, metadata, ctx)

}

func (c *PogodocClient) SaveTemplateFromFileStream(fsProps FileStreamProps, metadata api.SaveCreatedTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	response, err := c.Templates().InitializeTemplateCreation(ctx)
	if err != nil {
		fmt.Println("Error initializing template creation:", err)
		return "", err
	}
	templateId := response.JobId

	err = UploadToS3WithURL(response.PresignedTemplateUploadUrl, fsProps, "application/zip")
	if err != nil {
		fmt.Println("Error uploading to S3:", err)
		return "", err
	}

	req := map[string]any{}
	err = c.Templates().ExtractTemplateFiles(ctx, templateId, req)
	if err != nil {
		fmt.Println("Error extracting template files:", err)
		return "", err
	}
	request := api.GenerateTemplatePreviewsRequest{
		Type: api.GenerateTemplatePreviewsRequestType(metadata.Type),
		Data: metadata.SampleData,
	}

	previewResponse, err := c.Templates().GenerateTemplatePreviews(ctx, templateId, &request)
	if err != nil {
		fmt.Println("Error generating template previews:", err)
		return "", err
	}
	previewPng := previewResponse.PngPreview.JobId
	previewPdf := previewResponse.PdfPreview.JobId

	saveCreatedTemplateRequest := api.SaveCreatedTemplateRequest{
		TemplateInfo: &api.SaveCreatedTemplateRequestTemplateInfo{
			Title:       metadata.Title,
			Description: metadata.Description,
			Categories:  metadata.Categories,
			Type:        metadata.Type,
			SourceCode:  metadata.SourceCode,
			SampleData:  metadata.SampleData,
		},
		PreviewIds: &api.SaveCreatedTemplateRequestPreviewIds{
			PngJobId: previewPng,
			PdfJobId: previewPdf,
		},
	}

	err = c.Templates().SaveCreatedTemplate(ctx, templateId, &saveCreatedTemplateRequest)
	if err != nil {
		fmt.Println("Error saving created template:", err)
		return "", err
	}

	return templateId, nil

}

func (c *PogodocClient) UpdateTemplate(templateId string, filePath string, metadata api.UpdateTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	payload, err := ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	payloadLength := len(payload)
	fsProps := FileStreamProps{
		payload:       payload,
		payloadLength: payloadLength,
	}

	return c.UpdateTemplateFromFileStream(templateId, fsProps, metadata, ctx)

}

func (c *PogodocClient) UpdateTemplateFromFileStream(templateId string, fsProps FileStreamProps, metadata api.UpdateTemplateRequestTemplateInfo, ctx context.Context) (string, error) {
	response, err := c.Templates().InitializeTemplateCreation(ctx)
	if err != nil {
		fmt.Println("Error initializing template creation:", err)
		return "", err
	}
	contentId := response.JobId

	err = UploadToS3WithURL(response.PresignedTemplateUploadUrl, fsProps, "application/zip")
	if err != nil {
		fmt.Println("Error uploading to S3:", err)
		return "", err
	}

	req := map[string]any{}
	err = c.Templates().ExtractTemplateFiles(ctx, contentId, req)
	if err != nil {
		fmt.Println("Error extracting template files:", err)
		return "", err
	}

	request := api.GenerateTemplatePreviewsRequest{
		Type: api.GenerateTemplatePreviewsRequestType(metadata.Type),
		Data: metadata.SampleData,
	}
	previewResponse, err := c.Templates().GenerateTemplatePreviews(ctx, contentId, &request)
	if err != nil {
		fmt.Println("Error generating template previews:", err)
		return "", err
	}

	updateTemplateReq := &api.UpdateTemplateRequest{
		TemplateInfo: &api.UpdateTemplateRequestTemplateInfo{
			Title:       metadata.Title,
			Type:        metadata.Type,
			Description: metadata.Description,
			Categories:  metadata.Categories,
			SourceCode:  metadata.SourceCode,
			SampleData:  metadata.SampleData,
		},
		PreviewIds: &api.UpdateTemplateRequestPreviewIds{
			PngJobId: previewResponse.PngPreview.JobId,
			PdfJobId: previewResponse.PdfPreview.JobId,
		},
		ContentId: contentId,
	}

	_, err = c.Templates().UpdateTemplate(ctx, templateId, updateTemplateReq)
	if err != nil {
		fmt.Println("Error updating template:", err)
		return "", err
	}

	return templateId, nil

}

func (c *PogodocClient) GenerateDocument(gdProps GenerateDocumentProps, ctx context.Context) (*api.GetJobStatusResponse, error) {

	initRequest := gdProps.InitializeRenderJobRequest
	initResponse, err := c.Documents().InitializeRenderJob(ctx, &initRequest)
	if err != nil {
		fmt.Println("Error initializing render job:", err)
		return nil, err
	}

	Data := []byte(fmt.Sprint(gdProps.InitializeRenderJobRequest.Data))
	if initResponse != nil && initResponse.PresignedDataUploadUrl != nil {
		err = UploadToS3WithURL(*initResponse.PresignedDataUploadUrl, FileStreamProps{
			payload:       Data,
			payloadLength: len(Data),
		}, "application/json")
		if err != nil {
			fmt.Println("Error uploading data to S3:", err)
			return nil, err
		}
	}

	template := gdProps.template
	if template != nil && *template != "" && initResponse.PresignedTemplateUploadUrl != nil {
		err = UploadToS3WithURL(*initResponse.PresignedTemplateUploadUrl, FileStreamProps{
			payload:       []byte(*template),
			payloadLength: len(*template),
		}, "text/html")
		if err != nil {
			fmt.Println("Error uploading template to S3:", err)
			return nil, err

		}
	}

	_, err = c.Documents().StartRenderJob(
		ctx,
		initResponse.JobId,
		&gdProps.StartRenderJobRequest,
	)

	if err != nil {
		fmt.Println("Error starting render job:", err)
		return nil, err

	}

	result, err := c.Documents().GetJobStatus(ctx, initResponse.JobId)
	if err != nil {
		fmt.Println("Error getting job status:", err)
		return nil, err

	}

	return result, nil

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pogodocEnv := PogodocEnv{
		baseURL: os.Getenv("POGODOC_BASE_URL"),
		token:   os.Getenv("POGODOC_TOKEN"),
	}
	ctx := context.Background()
	c := PogodocClientInit(pogodocEnv.baseURL, pogodocEnv.token)

	response, err := c.Templates().InitializeTemplateCreation(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.JobId)
	sampledata, _ := ReadFile("../../data/json_data/react.json")

	var sampleDataMap map[string]interface{}
	err = json.Unmarshal([]byte(sampledata), &sampleDataMap)
	if err != nil {
		fmt.Println("Error unmarshaling sample data:", err)
		return
	}

	templateId, err := c.SaveTemplate("../../data/templates/React-Demo-App.zip", api.SaveCreatedTemplateRequestTemplateInfo{
		Title:       "Naslov",
		Description: "Deksripshn",
		Type:        api.SaveCreatedTemplateRequestTemplateInfoTypeReact,
		SampleData:  sampleDataMap,
		Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{1, 2},
	}, ctx)
	if err != nil {
		fmt.Println("Error saving template:", err)
		return
	}
	sourcecode := "SORSKODE"

	_, err = c.UpdateTemplate(templateId, "../../data/templates/React-Demo-App.zip", api.UpdateTemplateRequestTemplateInfo{
		Title:       "Naslov SMENET",
		Description: "ANDREJ UPDATE TEMPLATE",
		Type:        api.UpdateTemplateRequestTemplateInfoTypeReact,
		SampleData:  sampleDataMap,
		SourceCode:  &sourcecode,
		Categories:  []api.UpdateTemplateRequestTemplateInfoCategoriesItem{1, 2},
	}, ctx)
	if err != nil {
		fmt.Println("Error saving template:", err)
		return
	}

	fmt.Println(templateId)

	render, err := c.GenerateDocument(GenerateDocumentProps{
		InitializeRenderJobRequest: api.InitializeRenderJobRequest{
			TemplateId: &templateId,
			Type:       api.InitializeRenderJobRequestTypeReact,
			Target:     api.InitializeRenderJobRequestTargetPdf,
			Data:       sampleDataMap,
		},
		StartRenderJobRequest: api.StartRenderJobRequest{
			ShouldWaitForRenderCompletion: Pointer(true),
		}}, ctx)

	if err != nil {
		fmt.Println("Error generating document:", err)
		return
	}
	fmt.Println("Document generated successfully:", render)

}
