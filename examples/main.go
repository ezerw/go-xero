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
	client := xero.NewClient(
		oauth2.NewClient(ctx, tokenSource),
		xero.TenantID(TenantID),
	)

	opts := xero.InvoiceListOptions{
		InvoiceNumbers: "INV-0016,INV-0008", // comma-separated list of invoice numbers
	}

	// List Invoices
	invoices, err := listInvoices(ctx, client, &opts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, i := range invoices {
		fmt.Printf("Invoice #%s\n", i.InvoiceNumber)
	}

	// Update an invoice
	//i, err := updateInvoice(ctx, client)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("Updated invoice #%s", i.InvoiceNumber)
}
