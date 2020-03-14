package main

import (
	"context"

	"github.com/ezerw/go-xero"
)

var account = &xero.Account{
	Code: "12315",
	Name: "Some original name22",
	Type: "BANK",
}

// Get a list of accounts.
func listAccounts(ctx context.Context, client *xero.Client, opts *xero.AccountListOptions) ([]*xero.Account, error) {
	return client.Accounts.List(ctx, opts)
}

// Get one account using AccountID as identifier.
func getAccountByID(ctx context.Context, client *xero.Client, accountID string) (*xero.Account, error) {
	return client.Accounts.GetByID(ctx, accountID)
}

// Create one Account.
func createAccount(ctx context.Context, client *xero.Client) (*xero.Account, error) {
	return client.Accounts.Create(ctx, account)
}

// Update an Account.
func updateAccount(ctx context.Context, client *xero.Client) (*xero.Account, error) {
	// Update one field of the invoice
	account.Description = "Updated account"

	return client.Accounts.Update(ctx, account)
}
