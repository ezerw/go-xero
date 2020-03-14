package main

import (
	"context"
	"github.com/ezerw/go-xero"
)

var invoice = &xero.Invoice{
	InvoiceID:       "96988e67-ecf9-466d-bfbf-0afa1725a649",
	Type:            "ACCREC",
	Reference:       "Some reference",
	LineAmountTypes: "Inclusive",
	LineItems: []xero.LineItem{{
		Description: "Line Item #1",
		Quantity:    1.00,
		UnitAmount:  10.50,
		AccountCode: "200",
	}, {
		Description: "Line Item #2",
		Quantity:    3.50,
		UnitAmount:  27.30,
		AccountCode: "200",
	}},
}

// Get a list of invoices.
func listInvoices(ctx context.Context, client *xero.Client, opts *xero.InvoiceListOptions) ([]*xero.Invoice, error) {
	return client.Invoices.List(ctx, opts)
}

// Get one invoice by InvoiceID or InvoiceNumber identifier.
func getInvoiceByID(ctx context.Context, client *xero.Client, ID string) (*xero.Invoice, error) {
	return client.Invoices.GetByID(ctx, ID)
}

// Create one invoice.
func createInvoice(ctx context.Context, client *xero.Client) (*xero.Invoice, error) {
	return client.Invoices.Create(ctx, invoice)
}

// Update an Invoice.
func updateInvoice(ctx context.Context, client *xero.Client) (*xero.Invoice, error) {
	// Update one field of the invoice
	invoice.Reference = "New reference 69"

	return client.Invoices.Update(ctx, invoice)
}
