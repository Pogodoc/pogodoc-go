package pogodoc

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// UploadToS3WithURL uploads a file to a presigned URL on S3.
// It takes the presigned URL, file stream properties, and content type as parameters.
// The file stream properties include the payload (file content) and its length.
// It returns an error if the upload fails or if required headers are missing.
func UploadToS3WithURL(predsignedURL string, fsProps FileStreamProps, contentType string) error {

	headers := http.Header{}
	if contentType != "" {
		headers.Set("Content-Type", string(contentType))
	} else {
		return fmt.Errorf(" Content-Type is empty")
	}

	if fsProps.payloadLength > 0 {
		headers.Set("Content-Length", fmt.Sprintf("%d", fsProps.payloadLength))
	} else {
		return fmt.Errorf(" Content-Length is empty")
	}
	client := &http.Client{}

	req, err := http.NewRequest("PUT", predsignedURL, bytes.NewBuffer(fsProps.payload))
	if err != nil {
		return fmt.Errorf("creating request: %v", err)
	}
	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("uploading file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("uploading file: %s", resp.Status)
	}

	return nil
}

// ReadFile reads the content of a file from the given file path.
// It resolves the absolute path of the file, opens it, and reads its content.
// If the file is empty, it returns an error.
// It returns the file content as a byte slice or an error if any step fails.
func ReadFile(filePath string) ([]byte, error) {
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error resolving absolute path:", err)
		return nil, err
	}

	file, err := os.Open(absolutePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	payload, err := os.ReadFile(absolutePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	payloadLength := len(payload)
	if payloadLength == 0 {
		fmt.Println("Error: File is empty")
		return nil, fmt.Errorf("error: File is empty")
	}

	return payload, nil
}

// Pointer is a utility function that returns a pointer to the given value.
// It is a generic function that can take any type T and returns a pointer to T.
func Pointer[T any](d T) *T {
    return &d
}

