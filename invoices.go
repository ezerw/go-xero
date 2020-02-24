package xero

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type InvoicesService service

// Invoice holds an Invoice structure.
type Invoice struct {
	Type            string        `json:"Type" xml:"Type"`
	InvoiceID       string        `json:"InvoiceID,omitempty"`
	InvoiceNumber   string        `json:"InvoiceNumber,omitempty"`
	Reference       string        `json:"Reference,omitempty"`
	Payments        []Payment     `json:"payments,omitempty"`
	CreditNotes     []CreditNote  `json:"CreditNotes,omitempty"`
	Prepayments     []Prepayment  `json:"Prepayments,omitempty"`
	OverPayments    []OverPayment `json:"Overpayments,omitempty"`
	AmountDue       float64       `json:"AmountDue,omitempty"`
	AmountPaid      float64       `json:"AmountPaid,omitempty"`
	AmountCredited  float64       `json:"AmountCredited,omitempty"`
	CurrencyRate    float64       `json:"CurrencyRate"`
	IsDiscounted    bool          `json:"IsDiscounted"`
	HasAttachments  bool          `json:"HasAttachments,omitempty"`
	HasErrors       bool          `json:"HasErrors"`
	Contact         Contact       `json:"Contact"`
	DateString      string        `json:"DateString,omitempty"`
	Date            string        `json:"Date"`
	DueDateString   string        `json:"DueDateString,omitempty"`
	DueDate         string        `json:"DueDate"`
	Status          string        `json:"Status"`
	LineAmountTypes string        `json:"LineAmountTypes,omitempty"`
	LineItems       []LineItem    `json:"LineItems"`
	SubTotal        float64       `json:"SubTotal,omitempty"`
	TotalTax        float64       `json:"TotalTax,omitempty"`
	Total           float64       `json:"Total,omitempty"`
	UpdatedDateUTC  string        `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    string        `json:"CurrencyCode"`
}

// Invoices holds a normal response from the invoices endpoint.
type Invoices struct {
	Invoices []Invoice `json:"Invoices"`
}

// List will returns all invoices for the authenticated user.
func (s *InvoicesService) List(ctx context.Context) (*Invoices, error) {
	u := fmt.Sprintf("%s/%s", s.client.BaseURL, "Invoices")

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Xero-tenant-id", string(s.client.TenantID))

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: xero responded with %d status code", res.StatusCode)
	}

	var invoices *Invoices
	if err := json.NewDecoder(res.Body).Decode(&invoices); err != nil {
		return nil, fmt.Errorf("error: failed to decode json. %v", err)
	}

	return invoices, nil
}
