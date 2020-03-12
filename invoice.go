package xero

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var InvoicesBaseURL = fmt.Sprintf("%s/Invoices", baseURL)

type InvoicesService service

// Invoice is an Accounts Payable or Accounts Receivable document in a Xero organisation.
type Invoice struct {
	Type                string        `json:"Type"`
	Contact             Contact       `json:"Contact"`
	Date                string        `json:"Date,omitempty"`
	DueDate             string        `json:"DueDate,omitempty"`
	Status              string        `json:"Status,omitempty"`
	LineAmountTypes     string        `json:"LineAmountTypes,omitempty"`
	LineItems           []LineItem    `json:"LineItems"`
	SubTotal            float64       `json:"SubTotal,omitempty"`
	TotalTax            float64       `json:"TotalTax,omitempty"`
	Total               float64       `json:"Total,omitempty"`
	TotalDiscount       float64       `json:"TotalDiscount,omitempty"`
	UpdatedDateUTC      string        `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode        string        `json:"CurrencyCode,omitempty"`
	CurrencyRate        float64       `json:"CurrencyRate,omitempty"`
	InvoiceID           string        `json:"InvoiceID,omitempty"`
	InvoiceNumber       string        `json:"InvoiceNumber,omitempty"`
	Reference           string        `json:"Reference,omitempty"`
	BrandingThemeID     string        `json:"BrandingThemeID,omitempty"`
	Url                 string        `json:"Url,omitempty"`
	SentToContact       bool          `json:"SentToContact,omitempty"`
	ExpectedPaymentDate string        `json:"ExpectedPaymentDate,omitempty"`
	PlannedPaymentDate  string        `json:"PlannedPaymentDate,omitempty"`
	HasAttachments      bool          `json:"HasAttachments,omitempty"`
	Payments            []Payment     `json:"Payments,omitempty"`
	CreditNotes         []CreditNote  `json:"CreditNotes,omitempty"`
	Prepayments         []Prepayment  `json:"Prepayments,omitempty"`
	OverPayments        []OverPayment `json:"Overpayments,omitempty"`
	AmountDue           float64       `json:"AmountDue,omitempty"`
	AmountPaid          float64       `json:"AmountPaid,omitempty"`
	CISDeduction        string        `json:"CISDeduction,omitempty"`
	FullyPaidOnDate     string        `json:"FullyPaidOnDate,omitempty"`
	AmountCredited      float64       `json:"AmountCredited,omitempty"`
}

// Invoices holds a normal response from the invoices endpoint.
type Invoices struct {
	Invoices []*Invoice `json:"Invoices"`
}

// InvoiceListOptions specifies the optional parameters to Get Invoices
type InvoiceListOptions struct {
	InvoiceID      string `url:"InvoiceID,omitempty"`
	InvoiceNumber  string `url:"InvoiceNumber,omitempty"`
	IDs            string `url:"IDs,omitempty"`
	InvoiceNumbers string `url:"InvoiceNumbers,omitempty"`
	ContactIDs     string `url:"ContactIDs,omitempty"`
	Statuses       string `url:"Statuses,omitempty"`
	CreatedByMyApp string `url:"createdByMyApp,omitempty"`
	Page           string `url:"page,omitempty"`
}

// List fetch the full list of invoices adding optional params to the URL.
// https://developer.xero.com/documentation/api/invoices#page-invoice
func (s *InvoicesService) List(ctx context.Context, opts *InvoiceListOptions) ([]*Invoice, error) {

	u, err := addOptions(InvoicesBaseURL, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var i *Invoices
	_, err = s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, err
	}

	return i.Invoices, nil
}

// Create creates and invoice.
func (s *InvoicesService) Create(ctx context.Context, invoice *Invoice) (*Invoice, error) {
	req, err := s.client.NewRequest(http.MethodPut, InvoicesBaseURL, invoice)
	if err != nil {
		return nil, err
	}

	var i Invoices
	_, err = s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, err
	}

	return i.Invoices[0], nil
}

// CreateMulti creates multiple invoices on Xero in a single call.
func (s *InvoicesService) CreateMulti(ctx context.Context, invoices *Invoices) ([]*Invoice, error) {
	req, err := s.client.NewRequest(http.MethodPut, InvoicesBaseURL, invoices)
	if err != nil {
		return nil, err
	}

	var i *Invoices
	_, err = s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, err
	}

	return i.Invoices, nil
}

// Update updates an invoice in Xero.
// It receives an existing Xero invoice and does a POST request to update it using the existing InvoiceID.
func (s *InvoicesService) Update(ctx context.Context, invoice *Invoice) (*Invoice, error) {
	if invoice.InvoiceID == "" {
		return nil, errors.New("the InvoiceID field must not be empty")
	}

	if invoice.Type == "" {
		return nil, errors.New("the Type field must not be empty")
	}

	req, err := s.client.NewRequest(http.MethodPost, InvoicesBaseURL, invoice)
	if err != nil {
		return nil, err
	}

	var i Invoices
	_, err = s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, err
	}

	return i.Invoices[0], nil
}
