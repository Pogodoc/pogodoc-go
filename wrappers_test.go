package pogodoc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	api "github.com/pogodoc/pogodoc-go/sdk"

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
	templateId    string
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
	c, err := PogodocClientInit(pogodocEnv.baseURL, pogodocEnv.token)
	if err != nil {
		fmt.Errorf("Error initializing PogodocClient")
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
		Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
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
			Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
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
			Categories:  []api.UpdateTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
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
			Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{"invoice", "report"},
		},
		data.ctx,
	)
	fmt.Println(templateId)
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

	doc, err := data.client.GenerateDocument(documentProps, data.ctx)
	fmt.Println(doc.Output.Data.Url)
	if err != nil {
		t.Errorf("GenerateDocument failed: %v", err)

	}

}
