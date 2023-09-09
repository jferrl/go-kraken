package kraken

import (
	"context"
	"net/http"
	"net/url"
)

// WebsocketsAuth handles communication with the WebsocketsAuth related
// methods of the Kraken API.
type WebsocketsAuth service

func (ws *WebsocketsAuth) WebsocketsToken(ctx context.Context) (*WebsocketsToken, error) {
	req, err := ws.client.newPrivateRequest(ctx, http.MethodPost, "GetWebSocketsToken", url.Values{})
	if err != nil {
		return nil, err
	}

	var v WebsocketsToken
	if err := ws.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}