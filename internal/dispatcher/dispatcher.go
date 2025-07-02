package dispatcher

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mrtuuro/auto-messager/internal/apperror"
	"github.com/mrtuuro/auto-messager/internal/code"
)

type Dispatcher struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func New(url, key string) *Dispatcher {
	return &Dispatcher{
		baseURL: url,
		apiKey:  key,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (d *Dispatcher) Send(ctx context.Context, to, content string) (string, error) {
	body := map[string]string{"to": to, "content": content}
	b, _ := json.Marshal(body)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, d.baseURL, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ins-auth-key", d.apiKey)

	resp, err := d.client.Do(req)
	if err != nil {
		return "", apperror.NewAppError(
			code.ErrSystemInternal,
			err,
			code.GetErrorMessage(code.ErrSystemInternal),
			)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return "", apperror.NewAppError(
			"WEBHOOK_ERROR",
			fmt.Errorf("webhook returned %d", resp.StatusCode),
			"Unexpected status code returned.",
			)
	}
	var r struct {
		MessageID string `json:"messageId"`
	}
	return r.MessageID, json.NewDecoder(resp.Body).Decode(&r)
}
