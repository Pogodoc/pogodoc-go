package pogodoc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

type PogodocEnv struct {
	baseURL string `env:"LAMBDA_BASE_URL"`
	token   string `env:"POGODOC_API_TOKEN"`
}

type TestData struct {
	PogodocEnv    PogodocEnv
	client        PogodocClient
	ctx           context.Context
	sampleDataMap map[string]interface{}
}

func PrepareData() TestData {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pogodocEnv := PogodocEnv{
		baseURL: os.Getenv("LAMBDA_BASE_URL"),
		token:   os.Getenv("POGODOC_API_TOKEN"),
	}
	c, err := PogodocClientInit()
	if err != nil {
		fmt.Println("Error initializing PogodocClient")
		return TestData{}
	}
	ctx := context.Background()

	sampledata, _ := ReadFile("../../data/json_data/react.json")

	var sampleDataMap map[string]interface{}
	_ = json.Unmarshal([]byte(sampledata), &sampleDataMap)

	return TestData{
		PogodocEnv:    pogodocEnv,
		client:        *c,
		ctx:           ctx,
		sampleDataMap: sampleDataMap,
	}
}

func TestPogodocClient(t *testing.T) {
	data := PrepareData()
	_, err := PogodocClientInitWithConfig(data.PogodocEnv.baseURL, data.PogodocEnv.token)
	if err != nil {
		t.Errorf("PogodocClientInit failed: %v", err)
	}
}

func TestSaveTemplate(t *testing.T) {
	data := PrepareData()

	_, err := data.client.SaveTemplate("../../data/templates/React-Demo-App.zip", SaveCreatedTemplateRequestTemplateInfo{
		Title:       "Naslov",
		Description: "Deksripshn",
		Type:        SaveCreatedTemplateRequestTemplateInfoTypeReact,
		SampleData:  data.sampleDataMap,
		Categories:  []SaveCreatedTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
	}, data.ctx)
	if err != nil {
		t.Errorf("SaveTemplate failed: %v", err)
	}

}

func TestUpdateTemplate(t *testing.T) {
	data := PrepareData()
	templateId, err := data.client.SaveTemplate(
		"../../data/templates/React-Demo-App.zip",
		SaveCreatedTemplateRequestTemplateInfo{
			Title:       "Naslov",
			Description: "Deksripshn",
			Type:        SaveCreatedTemplateRequestTemplateInfoTypeReact,
			SampleData:  data.sampleDataMap,
			Categories:  []SaveCreatedTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
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
		UpdateTemplateRequestTemplateInfo{
			Title:       "Naslov SMENET",
			Description: "ANDREJ UPDATE TEMPLATE",
			Type:        UpdateTemplateRequestTemplateInfoTypeReact,
			SampleData:  data.sampleDataMap,
			SourceCode:  &src,
			Categories:  []UpdateTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
		},
		data.ctx,
	)
	if err != nil {
		t.Errorf("UpdateTemplate failed: %v", err)
	}
}

func TestGenerateDocument(t *testing.T) {
	data := PrepareData()

	sampleData := make(map[string]interface{})

	jsonData := `{
		"name": "John Doe"
	}`

	err := json.Unmarshal([]byte(jsonData), &sampleData)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}

	simpleDocumentProps := GenerateDocumentProps{
		InitializeRenderJobRequest: InitializeRenderJobRequest{
			TemplateId: String(os.Getenv("TEMPLATE_ID")),
			Type:       InitializeRenderJobRequestTypeHtml,
			Target:     InitializeRenderJobRequestTargetPdf,
			Data:       sampleData,
		},
	}

	startRenderJobResponse, err := data.client.StartGenerateDocument(simpleDocumentProps, data.ctx)
	fmt.Println("START GENERATE DOCUMENT: ", startRenderJobResponse.JobId)
	if err != nil {
		t.Errorf("GenerateDocument failed: %v", err)
	}

	generatedDocument, err := data.client.GenerateDocument(simpleDocumentProps, data.ctx)
	fmt.Println("GENERATE DOCUMENT: ", generatedDocument.Output.Data.Url)
	if err != nil {
		t.Errorf("GenerateDocument failed: %v", err)
	}

	immediateDocument, err := data.client.GenerateDocumentImmediate(simpleDocumentProps, data.ctx)
	fmt.Println("GENERATE DOCUMENT IMMEDIATE: ", immediateDocument.Url)
	if err != nil {
		t.Errorf("GenerateDocumentImmediate failed: %v", err)
	}
}

func TestReadMeExample(t *testing.T) {
	godotenv.Load()
	ctx := context.Background()

	client, err := PogodocClientInit()

	if err != nil {
		t.Errorf("PogodocClientInit failed: %v", err)
	}

	var sampleData map[string]interface{}

	jsonData := `{
		"name": "John Doe"
	}`

	err = json.Unmarshal([]byte(jsonData), &sampleData)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}

	documentProps := GenerateDocumentProps{
		InitializeRenderJobRequest: InitializeRenderJobRequest{
			TemplateId: String(os.Getenv("TEMPLATE_ID")),
			Type:       InitializeRenderJobRequestType("html"),
			Target:     InitializeRenderJobRequestTarget("pdf"),
			Data:       sampleData,
		},
	}

	doc, err := client.GenerateDocument(documentProps, ctx)
	if err != nil {
		t.Errorf("GenerateDocument failed: %v", err)
	}

	fmt.Println("README EXAMPLE: ", doc.Output.Data.Url)
}