package kraken

import (
	"net/http"
	"net/url"
)

const (
	defaultURL        = "https://api.kraken.com/"
	defaultAPIVersion = "0"

	userAgent = "go-kraken"
)

// Secrets represents the Kraken API key and secret.
type Secrets struct {
	Key    string
	Secret string
}

// A Client manages communication with the Mercedes API.
type Client struct {
	client *http.Client

	// Base URL for API requests. Defaults to the kraken API.
	//BaseURL should always be specified with a trailing slash.
	baseURL *url.URL

	apiKey APIKey // API key used for authentication.

	signer *Signer // Signer used to sign API requests.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Kraken API.
	Market         *MarketData
	Account        *Account
	Trading        *Trading
	WebsocketsAuth *WebsocketsAuth
}

type service struct {
	client *Client
}

// New returns new Kraken API client.If a nil httpClient is
// provided, a new http.Client will be used.
func New(httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(defaultURL + defaultAPIVersion + "/")

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		baseURL: baseURL,
		client:  httpClient,
	}

	c.common.client = c

	c.Market = (*MarketData)(&c.common)
	c.Account = (*Account)(&c.common)
	c.Trading = (*Trading)(&c.common)
	c.WebsocketsAuth = (*WebsocketsAuth)(&c.common)

	return c
}

// WithAuth sets the Kraken API key and secret.
func (c *Client) WithAuth(s Secrets) *Client {
	c.apiKey = APIKey(s.Key)
	c.signer = NewSigner(s.Secret)

	return c
}
