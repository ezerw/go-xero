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
		Description: "Single hammer",
		Quantity:    1.00,
		UnitAmount:  20.00,
		AccountCode: "200",
	}

	invoice := &xero.Invoice{
		Type:      "ACCREC",
		LineItems: []xero.LineItem{},
	}

	invoice.LineItems = append(invoice.LineItems, lineItem)

	updatedInvoice, err := xeroClient.Invoices.Update(ctx, "INVOICE_ID", invoice)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}

	fmt.Printf("%+v\n", updatedInvoice)
}
