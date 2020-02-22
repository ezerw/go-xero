package xero

import (
	"fmt"
	"net/http"
	"net/url"
)

const collectionURL = "Invoices"

type (
	// iClient is the Invoice client.
	iClient struct {
		Request *http.Request
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
	iUrl := fmt.Sprintf("%s/%s", ApiUrl, collectionURL)

	req, _ := http.NewRequest(http.MethodGet, iUrl, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Config.AccessToken))
	req.Header.Set("Xero-tenant-id", fmt.Sprintf("%s", c.Config.TenantID))
	req.Header.Set("Accept", "application/json")

	return &iClient{Request: req}
}

func (i *iClient) Get() ([]Invoice, error) {
	var invoices Invoices
	if err := getJson(i.Request, &invoices); err != nil {
		return nil, err
	}

	return invoices.Invoices, nil
}

func (i *iClient) Find(invoiceID string) (*Invoice, error) {
	iUrl, err := url.Parse(i.Request.URL.String() + "/" + invoiceID)
	if err != nil {
		return nil, err
	}
	i.Request.URL = iUrl

	var invoices Invoices
	if err := getJson(i.Request, &invoices); err != nil {
		return nil, err
	}

	return &invoices.Invoices[0], nil
}
