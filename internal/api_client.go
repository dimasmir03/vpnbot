package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{BaseURL: baseURL}
}

func (c *APIClient) SendRequest(endpoint string, method string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *APIClient) CreateUSer(user map[string]interface{}) (string, error) {
	res, err := c.SendRequest("users/create", "POST", user)
	if err != nil {
		return "", nil
	}
	return string(res), nil
}
