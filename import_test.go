package pogodoc

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestImportTypes(t *testing.T) {
	// Test struct type imports
	initRenderJobReq := InitializeRenderJobRequest{
		Type:   InitializeRenderJobRequestTypeHtml,
		Target: InitializeRenderJobRequestTargetPdf,
		Data:   map[string]interface{}{"name": "Test Document"},
	}

	startRenderJobReq := StartRenderJobRequest{
		ShouldWaitForRenderCompletion: Bool(true),
	}

	// Test using the pointer helpers
	assert.NotNil(t, Bool(true))
	assert.NotNil(t, String("test"))
	assert.NotNil(t, Int(42))
	assert.NotNil(t, Float64(3.14))
	assert.NotNil(t, Time(time.Now()))

	// Test enum constants
	assert.Equal(t, "pdf", string(InitializeRenderJobRequestTargetPdf))
	assert.Equal(t, "html", string(InitializeRenderJobRequestTypeHtml))

	// Test GenerateDocumentProps works correctly
	template := "<html><body>Hello {{name}}</body></html>"
	props := GenerateDocumentProps{
		InitializeRenderJobRequest: initRenderJobReq,
		StartRenderJobRequest:      startRenderJobReq,
		Template:                   &template,
	}

	assert.Equal(t, InitializeRenderJobRequestTypeHtml, props.InitializeRenderJobRequest.Type)
	assert.Equal(t, InitializeRenderJobRequestTargetPdf, props.InitializeRenderJobRequest.Target)
	assert.Equal(t, true, *props.StartRenderJobRequest.ShouldWaitForRenderCompletion)
	assert.Equal(t, "<html><body>Hello {{name}}</body></html>", *props.Template)

	// Test Environments works
	assert.Equal(t, "https://api.pogodoc.com/v1", Environments.Default)

	// Don't actually create a client unless POGODOC_API_TOKEN is set
	if os.Getenv("POGODOC_API_TOKEN") != "" {
		client, err := PogodocClientInit()
		assert.NoError(t, err)
		assert.NotNil(t, client)

		// Test client initialization works with token
		clientWithToken, err := PogodocClientInitWithToken("test-token")
		assert.NoError(t, err)
		assert.NotNil(t, clientWithToken)
	}

	// Test helper functions for converting strings to enums
	target, err := NewInitializeRenderJobRequestTargetFromString("pdf")
	assert.NoError(t, err)
	assert.Equal(t, InitializeRenderJobRequestTargetPdf, target)

	// Test FileStreamProps
	fsProps := FileStreamProps{
		payload:       []byte("test data"),
		payloadLength: 9,
	}
	assert.Equal(t, 9, fsProps.payloadLength)
	assert.Equal(t, []byte("test data"), fsProps.payload)

	// Make sure context works with the client methods (though we won't call them)
	_ = context.Background()
}
