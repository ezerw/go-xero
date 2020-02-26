package xero

import (
	"context"
	"fmt"
	"net/http"
)

var InvoicesBaseURL = fmt.Sprintf("%s/Invoices", baseURL)

type InvoicesService service

// Invoice is an Accounts Payable or Accounts Receivable document in a Xero organisation
type Invoice struct {
	Type                string         `json:"Type"`
	Contact             Contact        `json:"Contact"`
	Date                string         `json:"Date,omitempty"`
	DueDate             string         `json:"DueDate,omitempty"`
	Status              string         `json:"Status,omitempty"`
	LineAmountTypes     string         `json:"LineAmountTypes,omitempty"`
	LineItems           []LineItem     `json:"LineItems"`
	SubTotal            float64        `json:"SubTotal,omitempty"`
	TotalTax            float64        `json:"TotalTax,omitempty"`
	Total               float64        `json:"Total,omitempty"`
	TotalDiscount       float64        `json:"TotalDiscount,omitempty"`
	UpdatedDateUTC      string         `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode        string         `json:"CurrencyCode,omitempty"`
	CurrencyRate        float64        `json:"CurrencyRate,omitempty"`
	InvoiceID           string         `json:"InvoiceID,omitempty"`
	InvoiceNumber       string         `json:"InvoiceNumber,omitempty"`
	Reference           string         `json:"Reference,omitempty"`
	BrandingThemeID     string         `json:"BrandingThemeID,omitempty"`
	Url                 string         `json:"Url,omitempty"`
	SentToContact       bool           `json:"SentToContact,omitempty"`
	ExpectedPaymentDate string         `json:"ExpectedPaymentDate,omitempty"`
	PlannedPaymentDate  string         `json:"PlannedPaymentDate,omitempty"`
	HasAttachments      bool           `json:"HasAttachments,omitempty"`
	Payments            *[]Payment     `json:"Payments,omitempty"`
	CreditNotes         *[]CreditNote  `json:"CreditNotes,omitempty"`
	Prepayments         *[]Prepayment  `json:"Prepayments,omitempty"`
	OverPayments        *[]OverPayment `json:"Overpayments,omitempty"`
	AmountDue           float64        `json:"AmountDue,omitempty"`
	AmountPaid          float64        `json:"AmountPaid,omitempty"`
	CISDeduction        string         `json:"CISDeduction,omitempty"`
	FullyPaidOnDate     string         `json:"FullyPaidOnDate,omitempty"`
	AmountCredited      float64        `json:"AmountCredited,omitempty"`
}

// Invoices holds a normal response from the invoices endpoint.
type Invoices struct {
	Invoices []*Invoice `json:"Invoices"`
}

// List will returns all invoices for the authenticated user.
func (s *InvoicesService) List(ctx context.Context) ([]*Invoice, error) {
	req, err := s.client.NewRequest(http.MethodGet, InvoicesBaseURL, nil)
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

func (s *InvoicesService) GetByID(ctx context.Context, invoiceID string) (*Invoice, error) {
	u := fmt.Sprintf("%s/%s", InvoicesBaseURL, invoiceID)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var i *Invoices
	_, err = s.client.Do(ctx, req, &i)
	if err != nil {
		return nil, err
	}

	if len(i.Invoices) == 0 {
		return nil, fmt.Errorf("invoice not found")
	}

	return i.Invoices[0], nil
}

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
