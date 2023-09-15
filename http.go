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

// Response represents a Kraken response.
type Response struct {
	Errors []string `json:"error"`
	Result any      `json:"result"`
}

// Error builds a Kraken API error.
func (r *Response) Error() error {
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

func (c *Client) newPrivateRequest(ctx context.Context, method string, path string, body reqBody) (*http.Request, error) {
	reqURL := c.buildPrivateURL(path)

	if otp := OtpFromContext(ctx); otp != "" {
		body.withOtp(otp)
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL.String(), strings.NewReader(body.string()))
	if err != nil {
		return nil, err
	}

	signature := c.signer.Sign(body, reqURL.Path)

	req.Header.Set("API-Key", string(c.apiKey))
	req.Header.Set("API-Sign", signature)
	req.Header.Set("Content-Type", body.contentType())

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
	var res Response

	res.Result = v

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil && err != io.EOF {
		return err
	}

	return res.Error()
}

type reqBody interface {
	string() string
	nonce() string
	withOtp(otp Otp)
	contentType() string
}

type formURLEncodedBody struct {
	url.Values
}

func newFormURLEncodedBody(b url.Values) formURLEncodedBody {
	if b == nil {
		b = url.Values{}
	}

	b.Set(nonceKey, fmt.Sprintf("%d", time.Now().UnixNano()))
	return formURLEncodedBody{b}
}

func (b formURLEncodedBody) string() string {
	return b.Encode()
}

func (b formURLEncodedBody) withOtp(otp Otp) {
	b.Set("otp", string(otp))
}

func (b formURLEncodedBody) nonce() string {
	return b.Get(nonceKey)
}

func (b formURLEncodedBody) contentType() string {
	return "application/x-www-form-urlencoded; charset=utf-8"
}

type jsonMessage map[string]any

type jsonBody struct {
	jsonMessage
}

func newJSONBody(b any) (jsonBody, error) {
	d, err := json.Marshal(b)
	if err != nil {
		return jsonBody{}, err
	}

	var msg jsonMessage
	err = json.Unmarshal(d, &msg)
	if err != nil {
		return jsonBody{}, err
	}

	msg["nonce"] = fmt.Sprintf("%d", time.Now().UnixNano())
	return jsonBody{msg}, nil
}

func (b jsonBody) string() string {
	str, _ := json.Marshal(b.jsonMessage)
	return string(str)
}

func (b jsonBody) withOtp(otp Otp) {
	b.jsonMessage["otp"] = string(otp)
}

func (b jsonBody) nonce() string {
	if nonce, ok := b.jsonMessage["nonce"].(string); ok {
		return nonce
	}
	return ""
}

func (b jsonBody) contentType() string {
	return "application/json"
}
