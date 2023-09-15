package kraken

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Subaccounts handles communication with the Subaccounts related
// methods of the Kraken API.
type Subaccounts service

// CreateSubaccountOpts represents the parameters to create a Subaccount.
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

// TransferOpts represents the parameters to transfer funds.
type TransferOpts struct {
	Asset  string `url:"asset,omitempty"`
	Amount string `url:"amount,omitempty"`
	From   string `url:"from,omitempty"`
	To     string `url:"to,omitempty"`
}

// Valid returns true if the TransferOpts is valid.
func (o TransferOpts) Valid() bool {
	return o.Asset != "" && o.Amount != "" && o.From != "" && o.To != ""
}

// Transfer transfers funds to and from master and subaccounts.
// Note: AccountTransfer must be called by the master account.
// Docs: https://docs.kraken.com/rest/#tag/Subaccounts/operation/accountTransfer
func (s *Subaccounts) Transfer(ctx context.Context, opts TransferOpts) (*TransferResult, error) {
	if !opts.Valid() {
		return nil, errors.New("invalid options")
	}

	body, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newPrivateRequest(ctx, http.MethodPost, "AccountTransfer", body)
	if err != nil {
		return nil, err
	}

	var v TransferResult
	if err := s.client.do(req, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
