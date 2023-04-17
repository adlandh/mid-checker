package driven

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/adlandh/mid-checker/domain"
)

var _ domain.PassportChangedEventSender = (*WebhookEventSender)(nil)

type WebhookEventSender struct {
	url    string
	client *http.Client
}

func NewWebhookEventSender(client *http.Client, cfg *Config) *WebhookEventSender {
	return &WebhookEventSender{
		url:    cfg.WebhookURL,
		client: client,
	}
}

func (z WebhookEventSender) SendChangedStatus(info *domain.PassportInfo) error {
	queryParams := url.Values{}
	queryParams.Set("changed", "true")
	queryParams.Set("percent", strconv.Itoa(info.InternalStatus.Percent))
	queryParams.Set("status", info.InternalStatus.Name)

	resp, err := z.client.Get(z.url + "?" + queryParams.Encode())
	if err != nil {
		return fmt.Errorf("error connecting to webhook %s: %w", z.url, err)
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("error geting data from %s: status is %s", z.url, resp.Status)
	}

	return nil
}
