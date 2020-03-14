package xero

import (
	"context"
	"fmt"
	"net/http"
)

// AccountsBaseURL holds the accounts API base endpoint.
var AccountsBaseURL = fmt.Sprintf("%s/Accounts", baseURL)

// AccountsService is the service to talk to accounts endpoint in Xero.
type AccountsService service

const (
	// Account types
	// https://developer.xero.com/documentation/api/types#AccountTypes

	AccountTypeBank                    = "BANK"                    // Bank account
	AccountTypeCurrentAsset            = "CURRENT"                 // Current Asset account
	AccountTypeCurrentLiability        = "CURRLIAB"                // Current Liability account
	AccountTypeDepreciation            = "DEPRECIATN"              // Depreciation account
	AccountTypeDirectCosts             = "DIRECTCOSTS"             // Direct Costs account
	AccountTypeEquity                  = "EQUITY"                  // Equity account
	AccountTypeExpense                 = "EXPENSE"                 // Expense account
	AccountTypeFixedAsset              = "FIXED"                   // Fixed Asset account
	AccountTypeInventoryAsset          = "INVENTORY"               // Inventory Asset account
	AccountTypeLiability               = "LIABILITY"               // Liability account
	AccountTypeNonCurrentAsset         = "NONCURRENT"              // Non-current Asset account
	AccountTypeOtherIncome             = "OTHERINCOME"             // Other Income account
	AccountTypeOverhead                = "OVERHEADS"               // Overhead account
	AccountTypePrepayment              = "PREPAYMENT"              // Prepayment account
	AccountTypeRevenue                 = "REVENUE"                 // Revenue account
	AccountTypeSales                   = "SALES"                   // Sale account
	AccountTypeNonCurrentLiability     = "TERMLIAB"                // Non-current Liability account
	AccountTypePAYGLiability           = "PAYGLIABILITY"           // PAYG Liability account
	AccountTypeSuperannuationExpense   = "SUPERANNUATIONEXPENSE"   // Superannuation Expense account
	AccountTypeSuperannuationLiability = "SUPERANNUATIONLIABILITY" // Superannuation Liability account
	AccountTypeWagesExpense            = "WAGESEXPENSE"            // Wages Expense account

	// Bank Account Types
	BankAccountTypeBank       = "BANK"       // Bank account
	BankAccountTypeCreditCard = "CREDITCARD" // Credit card account
	BankAccountTypePayPal     = "PAYPAL"     // Paypal account

	// Account statuses
	AccountStatusActive   = "ACTIVE"
	AccountStatusArchived = "ARCHIVED"
)

type Account struct {
	AccountID               string `json:"AccountID,omitempty"`
	Code                    string `json:"Code,omitempty"`
	Name                    string `json:"Name,omitempty"`
	Status                  string `json:"Status,omitempty"`
	Type                    string `json:"Type,omitempty"`
	TaxType                 string `json:"TaxType,omitempty"`
	Class                   string `json:"Class,omitempty"`
	SystemAccount           string `json:"SystemAccount,omitempty"`
	EnablePaymentsToAccount bool   `json:"EnablePaymentsToAccount,omitempty"`
	ShowInExpenseClaims     bool   `json:"ShowInExpenseClaims,omitempty"`
	Description             string `json:"Description,omitempty"`
	BankAccountNumber       string `json:"BankAccountNumber,omitempty"`
	BankAccountType         string `json:"BankAccountType,omitempty"`
	CurrencyCode            string `json:"CurrencyCode,omitempty"`
	ReportingCode           string `json:"ReportingCode,omitempty"`
	ReportingCodeName       string `json:"ReportingCodeName,omitempty"`
	HasAttachments          bool   `json:"HasAttachments,omitempty"`
	UpdatedDateUTC          string `json:"UpdatedDateUTC,omitempty"`
	AddToWatchlist          bool   `json:"AddToWatchlist,omitempty"`
}

// Accounts contains a collection of Accounts
type Accounts struct {
	Accounts []*Account `json:"Accounts,omitempty"`
}

// AccountListOptions specifies the optional parameters to Get Invoices
type AccountListOptions struct {
	Order string `url:"order"`
}

// List fetch the full list of accounts adding optional params to the URL.
func (s *AccountsService) List(ctx context.Context, opts *AccountListOptions) ([]*Account, error) {
	u, err := addOptions(AccountsBaseURL, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var a *Accounts
	_, err = s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, err
	}

	return a.Accounts, nil
}

// GetByID fetch an Account by its AccountID from Xero and returns it.
func (s *AccountsService) GetByID(ctx context.Context, accountID string) (*Account, error) {
	if accountID == "" {
		return nil, ValidationNotEmptyError{Field: "accountID"}
	}
	u := fmt.Sprintf("%s/%s", AccountsBaseURL, accountID)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var a *Accounts
	_, err = s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, err
	}

	if len(a.Accounts) == 0 {
		return nil, ResourceNotFoundError{Resource: "account", ID: accountID}
	}

	return a.Accounts[0], nil
}

// Create method creates an Account resource in Xero
// Note that creating accounts with BankAccountType = CREDITCARD or PAYPAL will result
// in an account with BankAccountType = BANK.
func (s *AccountsService) Create(ctx context.Context, account *Account) (*Account, error) {
	req, err := s.client.NewRequest(http.MethodPut, AccountsBaseURL, account)
	if err != nil {
		return nil, err
	}

	var a Accounts
	_, err = s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, err
	}

	return a.Accounts[0], nil
}

// Update updates an account in Xero.
// It receives an existing Xero account and does a POST request to update it using the existing AccountID.
func (s *AccountsService) Update(ctx context.Context, account *Account) (*Account, error) {
	if account.AccountID == "" {
		return nil, ValidationNotEmptyError{Field: "AccountID"}
	}

	req, err := s.client.NewRequest(http.MethodPost, AccountsBaseURL, account)
	if err != nil {
		return nil, err
	}

	var a Accounts
	_, err = s.client.Do(ctx, req, &a)
	if err != nil {
		return nil, err
	}

	return a.Accounts[0], nil
}
