package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"

	"github.com/ezerw/xero"
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

	lineItem := xero.LineItem{
		Description: "Importing & Exporting Services",
		Quantity:    2.00,
		UnitAmount:  100.00,
		AccountCode: "200",
	}

	invoiceOne := &xero.Invoice{
		Type: "ACCREC",
		Contact: xero.Contact{
			Name: "Eze Rodriguez",
		},

		LineAmountTypes: "Inclusive",
		LineItems:       []xero.LineItem{},
	}

	invoiceTwo := &xero.Invoice{
		Type: "ACCREC",
		Contact: xero.Contact{
			Name: "Eze Rodriguez",
		},

		LineAmountTypes: "Inclusive",
		LineItems:       []xero.LineItem{},
	}

	invoiceOne.LineItems = append(invoiceOne.LineItems, lineItem)
	invoiceTwo.LineItems = append(invoiceTwo.LineItems, lineItem)

	var invoices xero.Invoices
	invoices.Invoices = append(invoices.Invoices, invoiceOne, invoiceTwo)

	createdInvoices, err := xeroClient.Invoices.CreateMulti(ctx, &invoices)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}

	fmt.Printf("%+v\n", createdInvoices)
}
