// Package pogodoc provides an SDK for interacting with the Pogodoc API,
// a document generation and template management service.
// This package allows users to create, update, and generate documents
// using predefined templates.
// Designed to integrate seamlessly with the Pogodoc platform,
// this SDK provides an easy-to-use interface for managing templates
// and documents within the service.
package pogodoc

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Pogodoc/pogodoc-go/client/client"
	"github.com/Pogodoc/pogodoc-go/client/option"
)

// PogodocClient is a client for interacting with the Pogodoc API.
// PogodocClientInit initializes a PogodocClient with the base URL and token from environment variables.
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

// PogodocClientInitWithConfig initializes a PogodocClient with a custom base URL and token.
func PogodocClientInitWithConfig(baseURL string, tokenString string) (*PogodocClient, error) {

	c := client.NewClient(
		option.WithToken(tokenString),
		option.WithBaseURL(baseURL),
	)

	return &PogodocClient{Client: c}, nil
}

// PogodocClientInitWithToken initializes a PogodocClient with only a token, using the default base URL.
func PogodocClientInitWithToken(tokenString string) (*PogodocClient, error) {
	c := client.NewClient(
		option.WithToken(tokenString),
	)
	return &PogodocClient{Client: c}, nil
}

// SaveTemplate is a method extension to SaveTeamplateFromFileStream to save a template from a file path to the Pogodoc service. 
// It wraps the SaveTemplateFromFileStream method.
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

// SaveTemplateFromFileStream is a method that allows saving a template from a file stream.
// It initializes the template creation, uploads the file to the Pogodoc service, extracts the template files,
// generates previews, and saves the template with the provided metadata.
// It returns the template ID or an error if any step fails.
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


// UpdateTemplate is a method extension to UpdateTemplateFromFileStream to update an existing template directly from a file path.
// It wraps the UpdateTemplateFromFileStream method.
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

// UpdateTemplateFromFileStream is a method that allows updating a template from a file stream.
// It initializes the template creation, uploads the file to the Pogodoc service, extracts the template files,
// generates previews, and updates the template with the provided metadata.
// It returns the template ID or an error if any step fails.
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


// GenerateDocument generates a document using the provided properties and context.
// It initializes a render job, uploads the necessary data and template, starts the render job,
// and retrieves the job status.
// It returns the job status response or an error if any step fails.
func (c *PogodocClient) StartGenerateDocument(gdProps GenerateDocumentProps, ctx context.Context) (*StartRenderJobResponse, error) {

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
			return nil, fmt.Errorf("uploading document: %v", err)
		}
	}

	template := gdProps.template

	if template != "" && initResponse.PresignedTemplateUploadUrl != nil {
		err = UploadToS3WithURL(*initResponse.PresignedTemplateUploadUrl, FileStreamProps{
			payload:       []byte(template),
			payloadLength: len(template),
		}, "text/html")
		if err != nil {
			return nil, fmt.Errorf("uploading document: %v", err)
		}
	}

	result, err := c.Documents.StartRenderJob(
		ctx,
		initResponse.JobId,
		&gdProps.StartRenderJobRequest,
	)
	if err != nil {
		return nil, fmt.Errorf("starting render: %v", err)
	}

	return result, nil

}


func (c *PogodocClient) GenerateDocument(gdProps GenerateDocumentProps, ctx context.Context) (*GetJobStatusResponse, error) {
	initResponse, err := c.StartGenerateDocument(gdProps, ctx)
	if err != nil {
		return nil, fmt.Errorf("starting document generation: %v", err)
	}

	return c.PollForJobCompletion(initResponse.JobId, ctx)
}

func (c *PogodocClient) GenerateDocumentImmediate(gdProps GenerateDocumentProps, ctx context.Context) (*StartImmediateRenderResponse, error) {

	return c.Documents.StartImmediateRender(ctx, &StartImmediateRenderRequest{
		Template: &gdProps.template,
		TemplateId: gdProps.InitializeRenderJobRequest.TemplateId,
		StartImmediateRenderRequestData: gdProps.InitializeRenderJobRequest.Data,
		Type: StartImmediateRenderRequestType(gdProps.InitializeRenderJobRequest.Type),
		Target: StartImmediateRenderRequestTarget(gdProps.InitializeRenderJobRequest.Target),
	})
}

func (c *PogodocClient) PollForJobCompletion(jobId string, ctx context.Context) (*GetJobStatusResponse, error) {
	maxAttempts := 60
	intervalMs := 500

	time.Sleep(1 * time.Second)

	for range maxAttempts {
		jobStatus, err := c.Documents.GetJobStatus(ctx, jobId)
		if err != nil {
			return nil, fmt.Errorf("getting job status: %v", err)
		}

		if *jobStatus.Status == "done" {
			return jobStatus, nil
		}
		time.Sleep(time.Duration(intervalMs) * time.Millisecond)
	}

	return nil, fmt.Errorf("job %s not found", jobId)
}