package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"

	"github.com/ezerw/go-xero"
)

const (
	AccessToken string = ""
	TenantID    string = ""
)

func main() {
	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: AccessToken})
	httpClient := oauth2.NewClient(ctx, tokenSource)

	xeroClient := xero.NewClient(httpClient, xero.TenantID(TenantID))

	// Update an invoice
	i, err := updateInvoice(ctx, xeroClient)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Updated invoice #%s", i.InvoiceNumber)
}

func updateInvoice(ctx context.Context, client *xero.Client) (*xero.Invoice, error) {
	invoice := &xero.Invoice{
		InvoiceID: "50a9f77e-caf4-40ef-9ab0-0c771accfdbf",
		Type:      "ACCREC",
		Contact: xero.Contact{
			Name: "New",
		},
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
	invoice.Reference = "New reference"

	i, err := client.Invoices.Update(ctx, invoice)
	if err != nil {
		return nil, err
	}

	// All good
	return i, nil
}
