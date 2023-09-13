package kraken

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Subaccount handles communication with the Subaccounts related
// methods of the Kraken API.
type Subaccounts service

type CreateSubaccountOpts struct {
	Username string `url:"username,omitempty"`
	Email    string `url:"email,omitempty"`
}

// Valid returns true if the CreateSubaccountOpts is valid.
func (o CreateSubaccountOpts) Valid() bool {
	return o.Username != "" && o.Email != ""
}

// Create creates a new trading subaccount.
// Docs: https://docs.kraken.com/rest/#tag/Subaccounts/operation/createSubaccount.
func (s *Subaccounts) Create(ctx context.Context, opts CreateSubaccountOpts) (bool, error) {
	if !opts.Valid() {
		return false, errors.New("invalid options")
	}

	body, err := query.Values(opts)
	if err != nil {
		return false, err
	}

	req, err := s.client.newPrivateRequest(ctx, http.MethodPost, "CreateSubaccount", body)
	if err != nil {
		return false, err
	}

	var v bool
	if err := s.client.do(req, &v); err != nil {
		return false, err
	}

	return v, nil
}
