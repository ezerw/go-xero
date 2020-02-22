## Xero Client

WIP - Xero Client using OAuth 2.0

Usage:
```
config := &xero.Config{
    AccessToken: "OAUTH2_ACCESS_TOKEN",
    TenantID:    "TENANT_ID",
    Timeout:     5 * time.Second, // Request timeout
}

xeroAPI, err := xero.NewClient(config)
if err != nil {
    fmt.Println(fmt.Errorf("error creating the client: %v", err))
}

// Get All invoices
invoices, err := xeroAPI.Invoices().Get()
if err != nil {
    fmt.Println(fmt.Errorf("error getting all invoices: %v", err))
}

// Find one invoice
invoice, err := xeroAPI.Invoices().Find("INVOICE_ID")
if err != nil {
    fmt.Println(fmt.Errorf("error getting invoice: %v", err))
}

fmt.Printf("%+v\n", invoices)
fmt.Printf("%+v\n", invoice)
```