package kraken

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
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

func (c *Client) newPublicRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	url := c.baseURL.String() + "public/" + path

	req, err := http.NewRequestWithContext(ctx, method, url, body)
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
