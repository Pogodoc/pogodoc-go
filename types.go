package pogodoc

import (
	api "github.com/Pogodoc/pogodoc-go/client"
	"github.com/Pogodoc/pogodoc-go/client/client"
	"github.com/Pogodoc/pogodoc-go/client/core"
	"github.com/Pogodoc/pogodoc-go/client/option"
)

// PogodocClient is an interface wrapper for the generated client.
type PogodocClient struct {
	*client.Client
}

// FileStreamProps is a struct that holds the properties for file streams.
// It contains the payload as a byte slice and its length.
// This struct is mainly used to pass file data when saving or updating templates.
type FileStreamProps struct {
	payload       []byte
	payloadLength int
}

// GenerateDocumentProps is a struct that holds the properties for generating documents.
type GenerateDocumentProps struct {
	InitializeRenderJobRequest InitializeRenderJobRequest
	StartRenderJobRequest      StartRenderJobRequest
	template                   *string
}

// Document Types
type InitializeRenderJobRequest = api.InitializeRenderJobRequest
type StartImmediateRenderRequest = api.StartImmediateRenderRequest
type StartRenderJobRequest = api.StartRenderJobRequest
type GetJobStatusResponse = api.GetJobStatusResponse
type GetJobStatusResponseOutput = api.GetJobStatusResponseOutput
type GetJobStatusResponseOutputData = api.GetJobStatusResponseOutputData
type GetJobStatusResponseOutputMetadata = api.GetJobStatusResponseOutputMetadata
type GetJobStatusResponseTarget = api.GetJobStatusResponseTarget
type GetJobStatusResponseType = api.GetJobStatusResponseType
type InitializeRenderJobRequestFormatOpts = api.InitializeRenderJobRequestFormatOpts
type InitializeRenderJobRequestFormatOptsFormat = api.InitializeRenderJobRequestFormatOptsFormat
type InitializeRenderJobRequestTarget = api.InitializeRenderJobRequestTarget
type InitializeRenderJobRequestType = api.InitializeRenderJobRequestType
type InitializeRenderJobResponse = api.InitializeRenderJobResponse
type InitializeRenderJobResponseTarget = api.InitializeRenderJobResponseTarget
type StartImmediateRenderRequestFormatOpts = api.StartImmediateRenderRequestFormatOpts
type StartImmediateRenderRequestFormatOptsFormat = api.StartImmediateRenderRequestFormatOptsFormat
type StartImmediateRenderRequestTarget = api.StartImmediateRenderRequestTarget
type StartImmediateRenderRequestType = api.StartImmediateRenderRequestType
type StartImmediateRenderResponse = api.StartImmediateRenderResponse
type StartRenderJobResponse = api.StartRenderJobResponse

// Template Types
type GenerateTemplatePreviewsRequest = api.GenerateTemplatePreviewsRequest
type SaveCreatedTemplateRequest = api.SaveCreatedTemplateRequest
type CloneTemplateResponse = api.CloneTemplateResponse
type DeleteTemplateResponse = api.DeleteTemplateResponse
type GeneratePresignedGetUrlResponse = api.GeneratePresignedGetUrlResponse
type GenerateTemplatePreviewsRequestFormatOpts = api.GenerateTemplatePreviewsRequestFormatOpts
type GenerateTemplatePreviewsRequestFormatOptsFormat = api.GenerateTemplatePreviewsRequestFormatOptsFormat
type GenerateTemplatePreviewsRequestType = api.GenerateTemplatePreviewsRequestType
type GenerateTemplatePreviewsResponse = api.GenerateTemplatePreviewsResponse
type GenerateTemplatePreviewsResponsePdfPreview = api.GenerateTemplatePreviewsResponsePdfPreview
type GenerateTemplatePreviewsResponsePngPreview = api.GenerateTemplatePreviewsResponsePngPreview
type GetTemplateIndexHtmlResponse = api.GetTemplateIndexHtmlResponse
type InitializeTemplateCreationResponse = api.InitializeTemplateCreationResponse
type SaveCreatedTemplateRequestPreviewIds = api.SaveCreatedTemplateRequestPreviewIds
type SaveCreatedTemplateRequestTemplateInfo = api.SaveCreatedTemplateRequestTemplateInfo
type SaveCreatedTemplateRequestTemplateInfoCategoriesItem = api.SaveCreatedTemplateRequestTemplateInfoCategoriesItem
type SaveCreatedTemplateRequestTemplateInfoType = api.SaveCreatedTemplateRequestTemplateInfoType
type UpdateTemplateRequestPreviewIds = api.UpdateTemplateRequestPreviewIds
type UpdateTemplateRequestTemplateInfo = api.UpdateTemplateRequestTemplateInfo
type UpdateTemplateRequestTemplateInfoCategoriesItem = api.UpdateTemplateRequestTemplateInfoCategoriesItem
type UpdateTemplateRequestTemplateInfoType = api.UpdateTemplateRequestTemplateInfoType
type UpdateTemplateResponse = api.UpdateTemplateResponse
type UpdateTemplateRequest = api.UpdateTemplateRequest
type UploadTemplateIndexHtmlRequest = api.UploadTemplateIndexHtmlRequest

// Core Types
type HTTPClient = core.HTTPClient
type RequestOption = option.RequestOption
type RequestOptions = core.RequestOptions
type BaseURLOption = core.BaseURLOption
type HTTPClientOption = core.HTTPClientOption
type HTTPHeaderOption = core.HTTPHeaderOption
type BodyPropertiesOption = core.BodyPropertiesOption
type QueryParametersOption = core.QueryParametersOption
type MaxAttemptsOption = core.MaxAttemptsOption
type TokenOption = core.TokenOption
type APIError = core.APIError
type FileParam = api.FileParam

// Environment constants
var Environments = api.Environments

// Pointer utility functions
var Bool = api.Bool
var Byte = api.Byte
var Complex64 = api.Complex64
var Complex128 = api.Complex128
var Float32 = api.Float32
var Float64 = api.Float64
var Int = api.Int
var Int8 = api.Int8
var Int16 = api.Int16
var Int32 = api.Int32
var Int64 = api.Int64
var Rune = api.Rune
var String = api.String
var Uint = api.Uint
var Uint8 = api.Uint8
var Uint16 = api.Uint16
var Uint32 = api.Uint32
var Uint64 = api.Uint64
var Uintptr = api.Uintptr
var UUID = api.UUID
var Time = api.Time
var MustParseDate = api.MustParseDate
var MustParseDateTime = api.MustParseDateTime

// Option functions
var WithBaseURL = option.WithBaseURL
var WithHTTPClient = option.WithHTTPClient
var WithHTTPHeader = option.WithHTTPHeader
var WithBodyProperties = option.WithBodyProperties
var WithQueryParameters = option.WithQueryParameters
var WithMaxAttempts = option.WithMaxAttempts
var WithToken = option.WithToken

// Constants for enums
// Document Target Constants
const (
	GetJobStatusResponseTargetPdf  = api.GetJobStatusResponseTargetPdf
	GetJobStatusResponseTargetHtml = api.GetJobStatusResponseTargetHtml
	GetJobStatusResponseTargetDocx = api.GetJobStatusResponseTargetDocx
	GetJobStatusResponseTargetXlsx = api.GetJobStatusResponseTargetXlsx
	GetJobStatusResponseTargetPptx = api.GetJobStatusResponseTargetPptx
	GetJobStatusResponseTargetPng  = api.GetJobStatusResponseTargetPng
	GetJobStatusResponseTargetJpg  = api.GetJobStatusResponseTargetJpg
)

// Document Type Constants
const (
	GetJobStatusResponseTypeDocx  = api.GetJobStatusResponseTypeDocx
	GetJobStatusResponseTypeXlsx  = api.GetJobStatusResponseTypeXlsx
	GetJobStatusResponseTypePptx  = api.GetJobStatusResponseTypePptx
	GetJobStatusResponseTypeEjs   = api.GetJobStatusResponseTypeEjs
	GetJobStatusResponseTypeHtml  = api.GetJobStatusResponseTypeHtml
	GetJobStatusResponseTypeLatex = api.GetJobStatusResponseTypeLatex
	GetJobStatusResponseTypeReact = api.GetJobStatusResponseTypeReact
)

// Initialize Render Job Format Constants
const (
	InitializeRenderJobRequestFormatOptsFormatLetter  = api.InitializeRenderJobRequestFormatOptsFormatLetter
	InitializeRenderJobRequestFormatOptsFormatLegal   = api.InitializeRenderJobRequestFormatOptsFormatLegal
	InitializeRenderJobRequestFormatOptsFormatTabloid = api.InitializeRenderJobRequestFormatOptsFormatTabloid
	InitializeRenderJobRequestFormatOptsFormatLedger  = api.InitializeRenderJobRequestFormatOptsFormatLedger
	InitializeRenderJobRequestFormatOptsFormatA0      = api.InitializeRenderJobRequestFormatOptsFormatA0
	InitializeRenderJobRequestFormatOptsFormatA1      = api.InitializeRenderJobRequestFormatOptsFormatA1
	InitializeRenderJobRequestFormatOptsFormatA2      = api.InitializeRenderJobRequestFormatOptsFormatA2
	InitializeRenderJobRequestFormatOptsFormatA3      = api.InitializeRenderJobRequestFormatOptsFormatA3
	InitializeRenderJobRequestFormatOptsFormatA4      = api.InitializeRenderJobRequestFormatOptsFormatA4
	InitializeRenderJobRequestFormatOptsFormatA5      = api.InitializeRenderJobRequestFormatOptsFormatA5
	InitializeRenderJobRequestFormatOptsFormatA6      = api.InitializeRenderJobRequestFormatOptsFormatA6
)

// Initialize Render Job Target Constants
const (
	InitializeRenderJobRequestTargetPdf  = api.InitializeRenderJobRequestTargetPdf
	InitializeRenderJobRequestTargetHtml = api.InitializeRenderJobRequestTargetHtml
	InitializeRenderJobRequestTargetDocx = api.InitializeRenderJobRequestTargetDocx
	InitializeRenderJobRequestTargetXlsx = api.InitializeRenderJobRequestTargetXlsx
	InitializeRenderJobRequestTargetPptx = api.InitializeRenderJobRequestTargetPptx
	InitializeRenderJobRequestTargetPng  = api.InitializeRenderJobRequestTargetPng
	InitializeRenderJobRequestTargetJpg  = api.InitializeRenderJobRequestTargetJpg
)

// Initialize Render Job Type Constants
const (
	InitializeRenderJobRequestTypeDocx  = api.InitializeRenderJobRequestTypeDocx
	InitializeRenderJobRequestTypeXlsx  = api.InitializeRenderJobRequestTypeXlsx
	InitializeRenderJobRequestTypePptx  = api.InitializeRenderJobRequestTypePptx
	InitializeRenderJobRequestTypeEjs   = api.InitializeRenderJobRequestTypeEjs
	InitializeRenderJobRequestTypeHtml  = api.InitializeRenderJobRequestTypeHtml
	InitializeRenderJobRequestTypeLatex = api.InitializeRenderJobRequestTypeLatex
	InitializeRenderJobRequestTypeReact = api.InitializeRenderJobRequestTypeReact
)

// Initialize Render Job Response Target Constants
const (
	InitializeRenderJobResponseTargetPdf  = api.InitializeRenderJobResponseTargetPdf
	InitializeRenderJobResponseTargetHtml = api.InitializeRenderJobResponseTargetHtml
	InitializeRenderJobResponseTargetDocx = api.InitializeRenderJobResponseTargetDocx
	InitializeRenderJobResponseTargetXlsx = api.InitializeRenderJobResponseTargetXlsx
	InitializeRenderJobResponseTargetPptx = api.InitializeRenderJobResponseTargetPptx
	InitializeRenderJobResponseTargetPng  = api.InitializeRenderJobResponseTargetPng
	InitializeRenderJobResponseTargetJpg  = api.InitializeRenderJobResponseTargetJpg
)

// Start Immediate Render Format Constants
const (
	StartImmediateRenderRequestFormatOptsFormatLetter  = api.StartImmediateRenderRequestFormatOptsFormatLetter
	StartImmediateRenderRequestFormatOptsFormatLegal   = api.StartImmediateRenderRequestFormatOptsFormatLegal
	StartImmediateRenderRequestFormatOptsFormatTabloid = api.StartImmediateRenderRequestFormatOptsFormatTabloid
	StartImmediateRenderRequestFormatOptsFormatLedger  = api.StartImmediateRenderRequestFormatOptsFormatLedger
	StartImmediateRenderRequestFormatOptsFormatA0      = api.StartImmediateRenderRequestFormatOptsFormatA0
	StartImmediateRenderRequestFormatOptsFormatA1      = api.StartImmediateRenderRequestFormatOptsFormatA1
	StartImmediateRenderRequestFormatOptsFormatA2      = api.StartImmediateRenderRequestFormatOptsFormatA2
	StartImmediateRenderRequestFormatOptsFormatA3      = api.StartImmediateRenderRequestFormatOptsFormatA3
	StartImmediateRenderRequestFormatOptsFormatA4      = api.StartImmediateRenderRequestFormatOptsFormatA4
	StartImmediateRenderRequestFormatOptsFormatA5      = api.StartImmediateRenderRequestFormatOptsFormatA5
	StartImmediateRenderRequestFormatOptsFormatA6      = api.StartImmediateRenderRequestFormatOptsFormatA6
)

// Start Immediate Render Target Constants
const (
	StartImmediateRenderRequestTargetPdf  = api.StartImmediateRenderRequestTargetPdf
	StartImmediateRenderRequestTargetHtml = api.StartImmediateRenderRequestTargetHtml
	StartImmediateRenderRequestTargetDocx = api.StartImmediateRenderRequestTargetDocx
	StartImmediateRenderRequestTargetXlsx = api.StartImmediateRenderRequestTargetXlsx
	StartImmediateRenderRequestTargetPptx = api.StartImmediateRenderRequestTargetPptx
	StartImmediateRenderRequestTargetPng  = api.StartImmediateRenderRequestTargetPng
	StartImmediateRenderRequestTargetJpg  = api.StartImmediateRenderRequestTargetJpg
)

// Start Immediate Render Type Constants
const (
	StartImmediateRenderRequestTypeDocx  = api.StartImmediateRenderRequestTypeDocx
	StartImmediateRenderRequestTypeXlsx  = api.StartImmediateRenderRequestTypeXlsx
	StartImmediateRenderRequestTypePptx  = api.StartImmediateRenderRequestTypePptx
	StartImmediateRenderRequestTypeEjs   = api.StartImmediateRenderRequestTypeEjs
	StartImmediateRenderRequestTypeHtml  = api.StartImmediateRenderRequestTypeHtml
	StartImmediateRenderRequestTypeLatex = api.StartImmediateRenderRequestTypeLatex
	StartImmediateRenderRequestTypeReact = api.StartImmediateRenderRequestTypeReact
)

// Template Preview Format Constants
const (
	GenerateTemplatePreviewsRequestFormatOptsFormatLetter  = api.GenerateTemplatePreviewsRequestFormatOptsFormatLetter
	GenerateTemplatePreviewsRequestFormatOptsFormatLegal   = api.GenerateTemplatePreviewsRequestFormatOptsFormatLegal
	GenerateTemplatePreviewsRequestFormatOptsFormatTabloid = api.GenerateTemplatePreviewsRequestFormatOptsFormatTabloid
	GenerateTemplatePreviewsRequestFormatOptsFormatLedger  = api.GenerateTemplatePreviewsRequestFormatOptsFormatLedger
	GenerateTemplatePreviewsRequestFormatOptsFormatA0      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA0
	GenerateTemplatePreviewsRequestFormatOptsFormatA1      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA1
	GenerateTemplatePreviewsRequestFormatOptsFormatA2      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA2
	GenerateTemplatePreviewsRequestFormatOptsFormatA3      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA3
	GenerateTemplatePreviewsRequestFormatOptsFormatA4      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA4
	GenerateTemplatePreviewsRequestFormatOptsFormatA5      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA5
	GenerateTemplatePreviewsRequestFormatOptsFormatA6      = api.GenerateTemplatePreviewsRequestFormatOptsFormatA6
)

// Template Preview Type Constants
const (
	GenerateTemplatePreviewsRequestTypeDocx  = api.GenerateTemplatePreviewsRequestTypeDocx
	GenerateTemplatePreviewsRequestTypeXlsx  = api.GenerateTemplatePreviewsRequestTypeXlsx
	GenerateTemplatePreviewsRequestTypePptx  = api.GenerateTemplatePreviewsRequestTypePptx
	GenerateTemplatePreviewsRequestTypeEjs   = api.GenerateTemplatePreviewsRequestTypeEjs
	GenerateTemplatePreviewsRequestTypeHtml  = api.GenerateTemplatePreviewsRequestTypeHtml
	GenerateTemplatePreviewsRequestTypeLatex = api.GenerateTemplatePreviewsRequestTypeLatex
	GenerateTemplatePreviewsRequestTypeReact = api.GenerateTemplatePreviewsRequestTypeReact
)

// Template Categories Constants
const (
	SaveCreatedTemplateRequestTemplateInfoCategoriesItemInvoice = api.SaveCreatedTemplateRequestTemplateInfoCategoriesItemInvoice
	SaveCreatedTemplateRequestTemplateInfoCategoriesItemMail    = api.SaveCreatedTemplateRequestTemplateInfoCategoriesItemMail
	SaveCreatedTemplateRequestTemplateInfoCategoriesItemReport  = api.SaveCreatedTemplateRequestTemplateInfoCategoriesItemReport
	SaveCreatedTemplateRequestTemplateInfoCategoriesItemCv      = api.SaveCreatedTemplateRequestTemplateInfoCategoriesItemCv
	SaveCreatedTemplateRequestTemplateInfoCategoriesItemOther   = api.SaveCreatedTemplateRequestTemplateInfoCategoriesItemOther
)

// Save Template Type Constants
const (
	SaveCreatedTemplateRequestTemplateInfoTypeDocx  = api.SaveCreatedTemplateRequestTemplateInfoTypeDocx
	SaveCreatedTemplateRequestTemplateInfoTypeXlsx  = api.SaveCreatedTemplateRequestTemplateInfoTypeXlsx
	SaveCreatedTemplateRequestTemplateInfoTypePptx  = api.SaveCreatedTemplateRequestTemplateInfoTypePptx
	SaveCreatedTemplateRequestTemplateInfoTypeEjs   = api.SaveCreatedTemplateRequestTemplateInfoTypeEjs
	SaveCreatedTemplateRequestTemplateInfoTypeHtml  = api.SaveCreatedTemplateRequestTemplateInfoTypeHtml
	SaveCreatedTemplateRequestTemplateInfoTypeLatex = api.SaveCreatedTemplateRequestTemplateInfoTypeLatex
	SaveCreatedTemplateRequestTemplateInfoTypeReact = api.SaveCreatedTemplateRequestTemplateInfoTypeReact
)

// Update Template Categories Constants
const (
	UpdateTemplateRequestTemplateInfoCategoriesItemInvoice = api.UpdateTemplateRequestTemplateInfoCategoriesItemInvoice
	UpdateTemplateRequestTemplateInfoCategoriesItemMail    = api.UpdateTemplateRequestTemplateInfoCategoriesItemMail
	UpdateTemplateRequestTemplateInfoCategoriesItemReport  = api.UpdateTemplateRequestTemplateInfoCategoriesItemReport
	UpdateTemplateRequestTemplateInfoCategoriesItemCv      = api.UpdateTemplateRequestTemplateInfoCategoriesItemCv
	UpdateTemplateRequestTemplateInfoCategoriesItemOther   = api.UpdateTemplateRequestTemplateInfoCategoriesItemOther
)

// Update Template Type Constants
const (
	UpdateTemplateRequestTemplateInfoTypeDocx  = api.UpdateTemplateRequestTemplateInfoTypeDocx
	UpdateTemplateRequestTemplateInfoTypeXlsx  = api.UpdateTemplateRequestTemplateInfoTypeXlsx
	UpdateTemplateRequestTemplateInfoTypePptx  = api.UpdateTemplateRequestTemplateInfoTypePptx
	UpdateTemplateRequestTemplateInfoTypeEjs   = api.UpdateTemplateRequestTemplateInfoTypeEjs
	UpdateTemplateRequestTemplateInfoTypeHtml  = api.UpdateTemplateRequestTemplateInfoTypeHtml
	UpdateTemplateRequestTemplateInfoTypeLatex = api.UpdateTemplateRequestTemplateInfoTypeLatex
	UpdateTemplateRequestTemplateInfoTypeReact = api.UpdateTemplateRequestTemplateInfoTypeReact
)

// Helper functions for creating types from strings
var NewGetJobStatusResponseTargetFromString = api.NewGetJobStatusResponseTargetFromString
var NewGetJobStatusResponseTypeFromString = api.NewGetJobStatusResponseTypeFromString
var NewInitializeRenderJobRequestFormatOptsFormatFromString = api.NewInitializeRenderJobRequestFormatOptsFormatFromString
var NewInitializeRenderJobRequestTargetFromString = api.NewInitializeRenderJobRequestTargetFromString
var NewInitializeRenderJobRequestTypeFromString = api.NewInitializeRenderJobRequestTypeFromString
var NewInitializeRenderJobResponseTargetFromString = api.NewInitializeRenderJobResponseTargetFromString
var NewStartImmediateRenderRequestFormatOptsFormatFromString = api.NewStartImmediateRenderRequestFormatOptsFormatFromString
var NewStartImmediateRenderRequestTargetFromString = api.NewStartImmediateRenderRequestTargetFromString
var NewStartImmediateRenderRequestTypeFromString = api.NewStartImmediateRenderRequestTypeFromString
var NewGenerateTemplatePreviewsRequestFormatOptsFormatFromString = api.NewGenerateTemplatePreviewsRequestFormatOptsFormatFromString
var NewGenerateTemplatePreviewsRequestTypeFromString = api.NewGenerateTemplatePreviewsRequestTypeFromString
var NewSaveCreatedTemplateRequestTemplateInfoCategoriesItemFromString = api.NewSaveCreatedTemplateRequestTemplateInfoCategoriesItemFromString
var NewSaveCreatedTemplateRequestTemplateInfoTypeFromString = api.NewSaveCreatedTemplateRequestTemplateInfoTypeFromString
var NewUpdateTemplateRequestTemplateInfoCategoriesItemFromString = api.NewUpdateTemplateRequestTemplateInfoCategoriesItemFromString
var NewUpdateTemplateRequestTemplateInfoTypeFromString = api.NewUpdateTemplateRequestTemplateInfoTypeFromString

