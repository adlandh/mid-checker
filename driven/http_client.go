package driven

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/ybbus/httpretry"
)

func NewHttpClient(cfg *Config) *http.Client {
	return httpretry.NewCustomClient(
		// Timeout *timeout
		&http.Client{Timeout: time.Duration(cfg.Retryer.Timeout) * time.Second, Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			TLSHandshakeTimeout: time.Duration(cfg.Retryer.TLSHandshakeTimeout) * time.Second,
		}},
		// retry up to *retries* times
		httpretry.WithMaxRetryCount(5),
		// retry on status >= 500, if err != nil, or if response was nil (status == 0)
		httpretry.WithRetryPolicy(func(statusCode int, err error) bool {
			return err != nil || statusCode >= 500 || statusCode == 0
		}),
		// every retry should wait one more second
		httpretry.WithBackoffPolicy(func(attemptNum int) time.Duration {
			return time.Duration(attemptNum+1) * time.Second
		}),
	)
}
