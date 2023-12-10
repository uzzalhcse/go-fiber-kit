// amd/client/client.go
package client

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

const BaseUrl = "https://test.api.amadeus.com"

type Client struct {
	apiKey      string
	apiSecret   string
	AccessToken string
	expiresAt   int64
	Client      *resty.Client
	mu          sync.Mutex
	BaseUrl     string
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		Client:    resty.New(),
		BaseUrl:   BaseUrl,
	}
}

func (c *Client) GetAccessToken() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.AccessToken != "" && time.Now().Unix() < c.expiresAt {
		return nil
	}

	resp, err := c.Client.R().
		SetFormData(map[string]string{
			"client_id":     c.apiKey,
			"client_secret": c.apiSecret,
			"grant_type":    "client_credentials",
		}).
		Post("https://test.api.amadeus.com/v1/security/oauth2/token")

	if err != nil {
		return err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return err
	}

	c.AccessToken = data["access_token"].(string)
	c.expiresAt = time.Now().Unix() + int64(data["expires_in"].(float64))
	return nil
}
