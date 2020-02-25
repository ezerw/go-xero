package xero

import (
	"context"
	"fmt"
	"net/http"
)

type InvoicesService service

// Invoice holds an Invoice structure.
type Invoice struct {
	Type                string        `json:"Type"`
	Contact             Contact       `json:"Contact"`
	Date                string        `json:"Date"`
	DueDate             string        `json:"DueDate"`
	Status              string        `json:"Status"`
	LineAmountTypes     string        `json:"LineAmountTypes,omitempty"`
	LineItems           []LineItem    `json:"LineItems"`
	SubTotal            float64       `json:"SubTotal,omitempty"`
	TotalTax            float64       `json:"TotalTax,omitempty"`
	Total               float64       `json:"Total,omitempty"`
	TotalDiscount       float64       `json:"TotalDiscount"`
	UpdatedDateUTC      string        `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode        string        `json:"CurrencyCode"`
	CurrencyRate        float64       `json:"CurrencyRate"`
	InvoiceID           string        `json:"InvoiceID"`
	InvoiceNumber       string        `json:"InvoiceNumber"`
	Reference           string        `json:"Reference,omitempty"`
	BrandingThemeID     string        `json:"BrandingThemeID,omitempty"`
	Url                 string        `json:"Url"`
	SentToContact       bool          `json:"SentToContact"`
	ExpectedPaymentDate string        `json:"ExpectedPaymentDate"`
	PlannedPaymentDate  string        `json:"PlannedPaymentDate"`
	HasAttachments      bool          `json:"HasAttachments"`
	Payments            []Payment     `json:"Payments,omitempty"`
	CreditNotes         []CreditNote  `json:"CreditNotes,omitempty"`
	Prepayments         []Prepayment  `json:"Prepayments,omitempty"`
	OverPayments        []OverPayment `json:"Overpayments,omitempty"`
	AmountDue           float64       `json:"AmountDue,omitempty"`
	AmountPaid          float64       `json:"AmountPaid,omitempty"`
	CISDeduction        string        `json:"CISDeduction"`
	FullyPaidOnDate     string        `json:"FullyPaidOnDate"`
	AmountCredited      float64       `json:"AmountCredited,omitempty"`
}

// Invoices holds a normal response from the invoices endpoint.
type Invoices struct {
	Invoices []*Invoice `json:"Invoices"`
}

// List will returns all invoices for the authenticated user.
func (s *InvoicesService) List(ctx context.Context) (*Invoices, error) {
	u := fmt.Sprintf("%s/%s", s.client.BaseURL, "Invoices")

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var invoices *Invoices
	_, err = s.client.Do(ctx, req, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (s *InvoicesService) GetByID(ctx context.Context, invoiceID string) (*Invoice, error) {
	u := fmt.Sprintf("%s/%s/%s", s.client.BaseURL, "Invoices", invoiceID)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var invoices *Invoices
	_, err = s.client.Do(ctx, req, &invoices)
	if err != nil {
		return nil, err
	}

	if len(invoices.Invoices) == 0 {
		return nil, fmt.Errorf("invoice not found")
	}

	return invoices.Invoices[0], nil

}
