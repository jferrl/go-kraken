package kraken

import (
	"context"
	"net/http"
)

// WebsocketsAuth handles communication with the WebsocketsAuth related
// methods of the Kraken API.
type WebsocketsAuth service

// WebsocketsToken retrieves a token to use for Websockets authentication.
func (ws *WebsocketsAuth) WebsocketsToken(ctx context.Context) (*WebsocketsToken, error) {
	req, err := ws.client.newPrivateRequest(ctx, http.MethodPost, "GetWebSocketsToken", newFormURLEncodedBody(nil))
	if err != nil {
		return nil, err
	}

	var v WebsocketsToken
	if err := ws.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
