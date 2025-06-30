package pogodoc

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

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

func Pointer[T any](d T) *T {
    return &d
}

