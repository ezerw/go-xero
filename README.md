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

### Get invoices
```go
invoices, err := xeroClient.Invoices.List(ctx, nil)
if err != nil {
    fmt.Println(fmt.Errorf("error getting invoices: %v", err))
}
```
To filter and paginate invoices you might use the same `List` method but passing pass an 
struct of type `InvoiceListOptions` as second parameter specifying those fields that will be included as query 
parameters part of the request URL
e.g:
```go
opts := InvoiceListOptions {
	InvoiceID: "",       // specific InvoiceID
	InvoiceNumber: "",   // specific InvoiceNumber
	IDs: "",             // comma-separated list of InvoicesIDs
	InvoiceNumbers: "",  // comma-separated list of InvoiceNumbers
	ContactIDs: "",      // comma-separated list of ContactIDs
	Statuses: "",        // comma-separated list of Statuses (DRAFT,SUBMITTED, etc)
	CreatedByMyApp: "",  // when set to true you'll only retrieve Invoices created by your app
	Page: "",            // Page:"1" – Up to 100 invoices will be returned in a single API call with line items shown for each invoice
}

filteredInvoices, err := xeroClient.Invoices.List(ctx, &opts)
...
```
This will always return a list of `0`, One or more `Invoices`  

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
// Existing Invoice
invoice := &xero.Invoice{
    InvoiceID: "EXISTING_INVOICE_ID",
    Type:      "ACCREC",
    Contact: xero.Contact{
        Name: "John Doe",
    },
    Reference:       "Invoice Reference",
    LineAmountTypes: "Inclusive",
    LineItems: []xero.LineItem{{
        Description: "Single Item",
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

fmt.Println("Invoice updated.")
fmt.Printf("InvoiceID: %s\n", i.InvoiceID)
fmt.Printf("InvoiceNumber: %s\n", i.InvoiceNumber)
fmt.Printf("Reference: %s\n", i.Reference)
```