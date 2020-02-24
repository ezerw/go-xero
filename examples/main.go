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

	invoices, err := xeroClient.Invoices.List(ctx)
	if err != nil {
		fmt.Println(fmt.Errorf("error getting invoices: %v", err))
	}

	fmt.Printf("%+v\n", invoices)
}
