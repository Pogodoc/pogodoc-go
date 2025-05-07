package main

import (
	"context"
	"fmt"
	"os"
	"pogodoc/go/sdk/client"
)

type PogodocClient struct {
	client.Client
}
func PogodocClientInit(baseURL string, tokenString string) PogodocClient {

	client := client.NewClient(
		client.ClientWithBaseURL(baseURL),
		client.ClientWithAuthToken(tokenString),
	)
	
	return PogodocClient{client} 

}
var Categories = map[string][]string{
	"categories": []string{"invoice", "mail", "report", "cv", "other"},
}

type TemplateFormatType string

const (
	Docx  TemplateFormatType = "docx"
	Xlsx  TemplateFormatType = "xlsx"
	Pptx  TemplateFormatType = "pptx"
	Ejs   TemplateFormatType = "ejs"
	Html  TemplateFormatType = "html"
	Latex TemplateFormatType = "latex"
	React TemplateFormatType = "react"
)
type TemplateMetadata struct {
	categories map[string][]string
	title string
	description string
	sampleData map[string]string
	fileType TemplateFormatType
	previewPng string
	previewPdf string
	sourceCode string
	
}


// Save Template function
func (c *PogodocClient)SaveTemplate(filePath string,ctx context.Context) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	payload, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	payloadLength := len(payload)
	if payloadLength == 0 {
		fmt.Println("Error: File is empty")
	}




	return ""
}


func main(){
	ctx := context.Background()
	client := PogodocClient("https://1mgxo5m49g.execute-api.eu-west-1.amazonaws.com/prod","eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklkIjoiNGQwMTdhOTEtODIwYy00N2FkLTlmMjAtZWQyYjg5YmJjYjcyIiwidXNlcklkIjoxLCJzdHJpcGVDdXN0b21lcklkIjoiY3VzX1JDS3EzNmx6bk5OQkNiIiwiaWF0IjoxNzQ1OTE1NzQwfQ.IZv4qQ3Y908Y1rso4ERJWQ3EqzMqjwFAOJrR0yjLMtSs4BknjVhODGDN2DUUDVHwJ28Pn2HYgCAM8gOPXSjspeN8mkIlWerbba6I3YsyAJUb2fdhmL2_vsqfIPfmA7EKe2W_V-YWRiQT4zFwsoD6iFgIGKhtvu2fyzoRUUDOtzdR5zgrmOLCjYwtolfPiw7xVsNNmoxrwdBAaW0esl9J3OTqspgeYZPwe35caOf9eQU3OxAyCmqd_gTGM3CJMPkHcyPDlM_qFZBl47jsviUAIxxxm0RmqVBfaebxobLPF2Zc3NgQF2ZZoMVmM1EVNnvv-45H4m3NfAzYdaf7NFCT-Q",)
	response, err := client.Templates().InitializeTemplateCreation(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.JobId)
}