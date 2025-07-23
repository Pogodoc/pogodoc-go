## Pogodoc Go SDK

The Pogodoc Go SDK enables developers to seamlessly generate documents and manage templates using Pogodoc’s API.

### Installation

To install the Go SDK, just execute the following command

```bash
$ go get github.com/Pogodoc/pogodoc-go
```

### Setup

To use the SDK you will need an API key which can be obtained from the [Pogodoc Dashboard](https://app.pogodoc.com)

### Example

```go

package main

import (
	"context"
	"encoding/json"
	"fmt"

	pogodoc "github.com/Pogodoc/pogodoc-go"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	ctx := context.Background()

	client, _ := pogodoc.PogodocClientInitWithToken("YOUR_POGODOC_API_TOKEN")

	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	var sampleData map[string]interface{}

	jsonData := `{
		"name": "John Doe"
	}`

	err = json.Unmarshal([]byte(jsonData), &sampleData)

	if err != nil {
		fmt.Println("Error unmarshalling JSON: %s", err)
		return
	}

	documentProps := pogodoc.GenerateDocumentProps{
		InitializeRenderJobRequest: pogodoc.InitializeRenderJobRequest{
			TemplateId: pogodoc.String("some-template-id"),
			Type:       pogodoc.InitializeRenderJobRequestType("ejs"),
			Target:     pogodoc.InitializeRenderJobRequestTarget("pdf"),
			Data:       sampleData,
		},
		StartRenderJobRequest: pogodoc.StartRenderJobRequest{
			ShouldWaitForRenderCompletion: pogodoc.Bool(true),
		}}

	doc, err := client.GenerateDocument(documentProps, ctx)

	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	fmt.Println(doc.Output.Data.Url)
}

```

### License

MIT License
