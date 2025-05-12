package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	//api "pogodoc/go/sdk"
	//"testing"
)

type TestData struct {
	client        PogodocClient
	ctx           context.Context
	templateId    string
	sampleDataMap map[string]interface{}
}

//	func main() {
//		err := godotenv.Load()
//		if err != nil {
//			log.Fatal("Error loading .env file")
//		}
//
//		pogodocEnv := PogodocEnv{
//			baseURL: os.Getenv("POGODOC_BASE_URL"),
//			token:   os.Getenv("POGODOC_TOKEN"),
//		}
//		ctx := context.Background()
//		c := PogodocClientInit(pogodocEnv.baseURL, pogodocEnv.token)
//
//		response, err := c.Templates().InitializeTemplateCreation(ctx)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(response.JobId)
//		sampledata, _ := ReadFile("../../data/json_data/react.json")
//
//		var sampleDataMap map[string]interface{}
//		err = json.Unmarshal([]byte(sampledata), &sampleDataMap)
//		if err != nil {
//			fmt.Println("Error unmarshaling sample data:", err)
//			return
//		}
//
//		templateId, err := c.SaveTemplate("../../data/templates/React-Demo-App.zip", api.SaveCreatedTemplateRequestTemplateInfo{
//			Title:       "Naslov",
//			Description: "Deksripshn",
//			Type:        api.SaveCreatedTemplateRequestTemplateInfoTypeReact,
//			SampleData:  sampleDataMap,
//			Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{1, 2},
//		}, ctx)
//		if err != nil {
//			fmt.Println("Error saving template:", err)
//			return
//		}
//		sourcecode := "SORSKODE"
//
//		_, err = c.UpdateTemplate(templateId, "../../data/templates/React-Demo-App.zip", api.UpdateTemplateRequestTemplateInfo{
//			Title:       "Naslov SMENET",
//			Description: "ANDREJ UPDATE TEMPLATE",
//			Type:        api.UpdateTemplateRequestTemplateInfoTypeReact,
//			SampleData:  sampleDataMap,
//			SourceCode:  &sourcecode,
//			Categories:  []api.UpdateTemplateRequestTemplateInfoCategoriesItem{1, 2},
//		}, ctx)
//		if err != nil {
//			fmt.Println("Error saving template:", err)
//			return
//		}
//
//		fmt.Println(templateId)
//
//		render, err := c.GenerateDocument(GenerateDocumentProps{
//			InitializeRenderJobRequest: api.InitializeRenderJobRequest{
//				TemplateId: &templateId,
//				Type:       api.InitializeRenderJobRequestTypeReact,
//				Target:     api.InitializeRenderJobRequestTargetPdf,
//				Data:       sampleDataMap,
//			},
//			StartRenderJobRequest: api.StartRenderJobRequest{
//				ShouldWaitForRenderCompletion: Pointer(true),
//			}}, ctx)
//
//		if err != nil {
//			fmt.Println("Error generating document:", err)
//			return
//		}
//		fmt.Println("Document generated successfully:", render)
//
// }
func PrepareData() TestData {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pogodocEnv := PogodocEnv{
		baseURL: os.Getenv("POGODOC_BASE_URL"),
		token:   os.Getenv("POGODOC_TOKEN"),
	}
	c := PogodocClientInit(pogodocEnv.baseURL, pogodocEnv.token)

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
		client:        c,
		ctx:           ctx,
		templateId:    response.JobId,
		sampleDataMap: sampleDataMap,
	}
}

//func TestPogodocClient(t *testing.T) {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//	pogodocEnv := PogodocEnv{
//		baseURL: os.Getenv("POGODOC_BASE_URL"),
//		token:   os.Getenv("POGODOC_TOKEN"),
//	}
//	PogodocClient := PogodocClientInit(pogodocEnv.baseURL, pogodocEnv.token)
//	if PogodocClient == nil {
//		t.Error("Failed to initialize PogodocClient")
//	}
//}
//
//func UpdateSaveTemplate(t *testing.T) {
//
//	testData := PrepareData()
//
//	c := testData.client
//	ctx := testData.ctx
//	templateId := testData.templateId
//	sampleDataMap := testData.sampleDataMap
//
//	_, err := c.SaveTemplate("../../data/templates/React-Demo-App.zip", api.SaveCreatedTemplateRequestTemplateInfo{
//		Title:       "Naslov",
//		Description: "Deksripshn",
//		Type:        api.SaveCreatedTemplateRequestTemplateInfoTypeReact,
//		SampleData:  sampleDataMap,
//		Categories:  []api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem{1, 2},
//	}, ctx)
//	if err != nil {
//		t.Error("Error saving template:", err)
//		return
//	}
//
//}
