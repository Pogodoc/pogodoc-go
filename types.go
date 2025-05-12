package main

import (
	api "pogodoc/go/sdk"
	"pogodoc/go/sdk/client"
)

type PogodocClient struct {
	client.Client
}
type FileStreamProps struct {
	payload       []byte
	payloadLength int
}

type GenerateDocumentProps struct {
	InitializeRenderJobRequest api.InitializeRenderJobRequest
	StartRenderJobRequest      api.StartRenderJobRequest
	template                   *string
}
