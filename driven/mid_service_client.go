package driven

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adlandh/mid-checker/domain"
)

var _ domain.PassportInfoFetcher = (*MidServiceClient)(nil)

type MidServiceClient struct {
	httpClient *http.Client
	Url        string
}

func NewMidServiceClient(client *http.Client, cfg *Config) *MidServiceClient {
	return &MidServiceClient{
		httpClient: client,
		Url:        "https://info.midpass.ru/api/request/" + cfg.PassportFormID,
	}
}

func (c *MidServiceClient) GetPassportStatus() (*domain.PassportInfo, error) {
	resp, err := c.httpClient.Get(c.Url)
	if err != nil {
		return nil, fmt.Errorf("error geting data from %s: %w", c.Url, err)
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("error geting data from %s: status is %s", c.Url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error geting request body from %s: %w", c.Url, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var response domain.PassportInfo

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal request body from %s: %w", c.Url, err)
	}

	return &response, nil
}
