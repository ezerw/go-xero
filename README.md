## Xero Client

WIP - Xero Client using OAuth 2.0

Usage:
```go
tokenSource := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "ACCESS_TOKEN"}
)

ctx := context.Background()
httpClient := oauth2.NewClient(ctx, tokenSource)

xeroClient := xero.NewClient(httpClient, "XERO_TENANT_ID")

invoices, err := xeroClient.Invoices.List(ctx)
if err != nil {
    fmt.Println(fmt.Errorf("error getting invoices: %v", err))
}

fmt.Printf("%+v\n", invoices)
```