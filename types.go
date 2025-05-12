package main

import (
	"pogodoc/go/sdk/client"
)

type PogodocClient struct {
	client.Client
}

// type SaveTemplateMetadata struct {
// 	TemplateInfo *api.SaveCreatedTemplateRequestTemplateInfo 			`json:"templateInfo,omitempty"`
// 	FormatOpts *api.GenerateTemplatePreviewsRequestFormatOpts           `json:"formatOpts,omitempty"`
// }

// type FormatOptsType struct {
// 	AdditionalProps map[string]interface{} `json:"-"`
// 	FromPage        *int                   `json:"fromPage,omitempty"`
// 	ToPage          *int                   `json:"toPage,omitempty"`
// 	Format          *string                `json:"format,omitempty"`
// 	WaitForSelector *string                `json:"waitForSelector,omitempty"`
// }

type FileStreamProps struct {

	payload []byte
	payloadLength int

}