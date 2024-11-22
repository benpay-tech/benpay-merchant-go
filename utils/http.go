package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// SendPostRequest send POST request
func SendPostRequest(url string, body []byte, headers map[string]string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, fmt.Errorf("create request err: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("send request err: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, fmt.Errorf("read response err: %w", err)
	}

	return resp, respBody, nil
}
