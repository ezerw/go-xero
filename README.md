# Xero Client

WIP - Xero Client using OAuth 2.0

## Usage:

```go
tokenSource := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "ACCESS_TOKEN"}
)

ctx := context.Background()
httpClient := oauth2.NewClient(ctx, tokenSource)

xeroClient := xero.NewClient(httpClient, "XERO_TENANT_ID")
````

### Get tenant invoices
```go
invoices, err := xeroClient.Invoices.List(ctx)
if err != nil {
    fmt.Println(fmt.Errorf("error getting invoices: %v", err))
}
```
### Get an Invoice by ID
```go
invoices, err := xeroClient.Invoices.GetById(ctx, "INVOICE_ID")
if err != nil {
    fmt.Println(fmt.Errorf("error getting invoice: %v", err))
}
```
### Create an Invoice
```go
lineItem := xero.LineItem{
    Description: "Importing & Exporting Services",
    Quantity:    2.00,
    UnitAmount:  100.00,
    AccountCode: "200",
}

invoice := &xero.Invoice{
    Type: "ACCREC",
    Contact: xero.Contact{
        Name: "Eze Rodriguez",
    },
    LineAmountTypes: "Inclusive",
    LineItems:       []xero.LineItem{},
}

invoice.LineItems = append(invoice.LineItems, lineItem)

cInvoice, err := xeroClient.Invoices.Create(ctx, invoice)
if err != nil {
    fmt.Println(fmt.Errorf("error: %v", err))
    return
}
```

### Create multiple invoices in one request
```go
lineItem := xero.LineItem{
    Description: "Services",
    Quantity:    1.00,
    UnitAmount:  55.00,
    AccountCode: "200",
}

invoiceOne := &xero.Invoice{
    Type: "ACCREC",
    Contact: xero.Contact{
        Name: "John Doe",
    },
    LineAmountTypes: "Inclusive",
    LineItems:       []xero.LineItem{},
}

invoiceTwo := &xero.Invoice{
    Type: "ACCREC",
    Contact: xero.Contact{
        Name: "Jane Doe",
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
```

### Update an Invoice
```go
lineItem := xero.LineItem{
    Description: "Single hammer",
    Quantity:    1.00,
    UnitAmount:  20.00,
    AccountCode: "200",
}

invoice := &xero.Invoice{
    Type:      "ACCREC", // Required by Xero API
    LineItems: []xero.LineItem{},
}

invoice.LineItems = append(invoice.LineItems, lineItem)

updatedInvoice, err := xeroClient.Invoices.Update(ctx, "INVOICE_ID", invoice)
if err != nil {
    fmt.Println(fmt.Errorf("error: %v", err))
    return
}
```