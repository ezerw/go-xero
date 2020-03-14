package main

import (
	"context"
	"github.com/ezerw/go-xero"
)

func listInvoices(ctx context.Context, client *xero.Client, opts *xero.InvoiceListOptions) ([]*xero.Invoice, error) {
	return client.Invoices.List(ctx, opts)
}

func updateInvoice(ctx context.Context, client *xero.Client) (*xero.Invoice, error) {
	// Update an invoice
	invoice := &xero.Invoice{
		InvoiceID: "50a9f77e-caf4-40ef-9ab0-0c771accfdbf",
		Type:      "ACCREC",
		Reference:       "Some reference",
		LineAmountTypes: "Inclusive",
		LineItems: []xero.LineItem{{
			Description: "Line Item #2",
			Quantity:    1.00,
			UnitAmount:  10.50,
			AccountCode: "200",
		}},
	}

	// Update one field of the invoice
	invoice.Reference = "New reference 45"

	return client.Invoices.Update(ctx, invoice)
}

func createInvoice(ctx context.Context, client *xero.Client) (*xero.Invoice, error) {
	invoice := &xero.Invoice{
		LineItems: []xero.LineItem{{
			Description: "Line Item #2",
			Quantity:    1.00,
			UnitAmount:  10.50,
			AccountCode: "200",
		}},
	}

	return client.Invoices.Create(ctx, invoice)
}
