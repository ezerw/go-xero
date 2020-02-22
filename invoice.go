package xero

import (
	"fmt"
	"net/http"
)

const collectionURL = "Invoices"

type (
	// iClient is the Invoice client.
	iClient struct {
		InvoicesBaseURL string
		AccessToken     AccessToken
		TenantID        TenantID
	}
	// Invoice holds an Invoice structure.
	Invoice struct {
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
	// Invoices holds a normal response to the invoices endpoint.
	Invoices struct {
		Invoices []Invoice `json:"Invoices"`
	}
)

func (c client) Invoices() *iClient {
	return &iClient{
		InvoicesBaseURL: fmt.Sprintf("%s/%s", ApiUrl, collectionURL),
		AccessToken:     c.Config.AccessToken,
		TenantID:        c.Config.TenantID,
	}
}

func (c *iClient) Get() ([]Invoice, error) {
	req, _ := http.NewRequest(http.MethodGet, c.InvoicesBaseURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Xero-tenant-id", fmt.Sprintf("%s", c.TenantID))
	req.Header.Set("Accept", "application/json")

	var i Invoices
	if err := getJson(req, &i); err != nil {
		return nil, err
	}

	return i.Invoices, nil
}

func (c *iClient) Find(invoiceID string) (*Invoice, error) {
	req, _ := http.NewRequest(http.MethodGet, c.InvoicesBaseURL+"/"+invoiceID, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.Header.Set("Xero-tenant-id", fmt.Sprintf("%s", c.TenantID))
	req.Header.Set("Accept", "application/json")

	var i Invoices
	if err := getJson(req, &i); err != nil {
		return nil, err
	}

	if i.Invoices == nil {
		return nil, fmt.Errorf("error: Invoice not found")
	}

	return &i.Invoices[0], nil
}

func (c *iClient) Create(i *Invoice) (*Invoice, error) {
}

func (c *iClient) Update(i *Invoice, ui Invoice) (*Invoice, error) {
	return nil, nil
}

func (c *iClient) Delete(i Invoice) error {
}
