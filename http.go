package kraken

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Resource represents a Kraken response.
type Resource struct {
	Errors []string `json:"error"`
	Result any      `json:"result"`
}

// Error builds a Kraken API error.
func (r *Resource) Error() error {
	if len(r.Errors) == 0 {
		return nil
	}

	return &Error{
		errors: r.Errors,
	}
}

func (c *Client) buildPublicURL(path string) *url.URL {
	u, _ := url.Parse(fmt.Sprintf("%spublic/%s", c.baseURL.String(), path))
	return u
}

func (c *Client) buildPrivateURL(path string) *url.URL {
	u, _ := url.Parse(fmt.Sprintf("%sprivate/%s", c.baseURL.String(), path))
	return u
}

func (c *Client) newPublicRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	reqURL := c.buildPublicURL(path).String()

	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *Client) newPrivateRequest(ctx context.Context, method string, path string, body url.Values) (*http.Request, error) {
	if body == nil {
		body = url.Values{}
	}

	reqURL := c.buildPrivateURL(path)

	if otp := OtpFromContext(ctx); otp != "" {
		body.Set("otp", string(otp))
	}

	body.Set(nonceKey, fmt.Sprintf("%d", time.Now().UnixNano()))

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	signature := c.signer.Sign(body, reqURL.Path)

	req.Header.Set("API-Key", string(c.apiKey))
	req.Header.Set("API-Sign", signature)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	return req, nil
}

func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		err = decodeResponse(resp, v)
	}

	return err
}

func decodeResponse(r *http.Response, v any) error {
	var resource Resource

	resource.Result = v

	err := json.NewDecoder(r.Body).Decode(&resource)
	if err != nil && err != io.EOF {
		return err
	}

	return resource.Error()
}
