package pogodoc

import (
	api "github.com/pogodoc/pogodoc-go/sdk"
	"github.com/pogodoc/pogodoc-go/sdk/client"
)

type PogodocClient struct {
	*client.Client
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
