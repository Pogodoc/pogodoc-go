package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	api "pogodoc/go/sdk"
	"testing"
)

type TestData struct {
	PogodocEnv    PogodocEnv
	client        PogodocClient
	ctx           context.Context
	templateId    string
	sampleDataMap map[string]interface{}
}

func PrepareData() TestData {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pogodocEnv := PogodocEnv{
		baseURL: os.Getenv("POGODOC_BASE_URL"),
		token:   os.Getenv("POGODOC_TOKEN"),
	}
	c, err := PogodocClientInit(pogodocEnv.baseURL, pogodocEnv.token)

	ctx := context.Background()
	response, err := c.Templates().InitializeTemplateCreation(ctx)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(response.JobId)
	sampledata, _ := ReadFile("../../data/json_data/react.json")

	var sampleDataMap map[string]interface{}
	_ = json.Unmarshal([]byte(sampledata), &sampleDataMap)

	return TestData{
		PogodocEnv:    pogodocEnv,
		client:        *c,
		ctx:           ctx,
		templateId:    response.JobId,
		sampleDataMap: sampleDataMap,
	}
}

func TestPogodocClient(t *testing.T) {
	data := PrepareData()
	_, err := PogodocClientInit(data.PogodocEnv.baseURL, data.PogodocEnv.token)
	if err != nil {
		t.Errorf("PogodocClientInit failed: %v", err)
	}
}

func TestSaveTemplate(t *testing.T) {
	data := PrepareData()
	_, err := data.client.SaveTemplate("../../data/templates/React-Demo-App.zip", api.SaveCreatedTemplateRequestTemplateInfo{
		Title:       "Naslov",
		Description: "Deksripshn",
		Type:        api.SaveCreatedTemplateRequestTemplateInfoTypeReact,
		SampleData:  data.sampleDataMap,
		Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{1, 2},
	}, data.ctx)
	if err != nil {
		t.Errorf("SaveTemplate failed: %v", err)
	}
}

func TestUpdateTemplate(t *testing.T) {
	data := PrepareData()
	templateId, err := data.client.SaveTemplate(
		"../../data/templates/React-Demo-App.zip",
		api.SaveCreatedTemplateRequestTemplateInfo{
			Title:       "Naslov",
			Description: "Deksripshn",
			Type:        api.SaveCreatedTemplateRequestTemplateInfoTypeReact,
			SampleData:  data.sampleDataMap,
			Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{1, 2},
		},
		data.ctx,
	)
	if err != nil {
		t.Fatalf("SaveTemplate failed: %v", err)
	}

	src := "SORSKODE"
	_, err = data.client.UpdateTemplate(
		templateId,
		"../../data/templates/React-Demo-App.zip",
		api.UpdateTemplateRequestTemplateInfo{
			Title:       "Naslov SMENET",
			Description: "ANDREJ UPDATE TEMPLATE",
			Type:        api.UpdateTemplateRequestTemplateInfoTypeReact,
			SampleData:  data.sampleDataMap,
			SourceCode:  &src,
			Categories:  []api.UpdateTemplateRequestTemplateInfoCategoriesItem{1, 2},
		},
		data.ctx,
	)
	if err != nil {
		t.Errorf("UpdateTemplate failed: %v", err)
	}
}

func TestGenerateDocument(t *testing.T) {
	data := PrepareData()

	templateId, err := data.client.SaveTemplate(
		"../../data/templates/React-Demo-App.zip",
		api.SaveCreatedTemplateRequestTemplateInfo{
			Title:       "Naslov",
			Description: "Deksripshn",
			Type:        api.SaveCreatedTemplateRequestTemplateInfoTypeReact,
			SampleData:  data.sampleDataMap,
			Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{1, 2},
		},
		data.ctx,
	)
	if err != nil {
		t.Fatalf("SaveTemplate failed: %v", err)
	}

	documentProps := GenerateDocumentProps{
		InitializeRenderJobRequest: api.InitializeRenderJobRequest{
			TemplateId: &templateId,
			Type:       api.InitializeRenderJobRequestTypeReact,
			Target:     api.InitializeRenderJobRequestTargetPdf,
			Data:       data.sampleDataMap,
		},
		StartRenderJobRequest: api.StartRenderJobRequest{
			ShouldWaitForRenderCompletion: Pointer(true),
		}}

	_, err = data.client.GenerateDocument(documentProps, data.ctx)
	if err != nil {
		t.Errorf("GenerateDocument failed: %v", err)

	}

}
