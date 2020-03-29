package xero

import (
	"testing"
)

func TestInvoice_marshall(t *testing.T) {
	testJSONMarshal(t, &Invoice{}, "{}")

	i := &Invoice{
		Type: "ACCREC",
		Contact: &Contact{
			ContactID:     "025867f1-d741-4d6b-b1af-9ac774b59ba7",
			ContactStatus: "ACTIVE",
			Name:          "City Agency",
			Addresses: []Address{
				{
					AddressType: "STREET",
				},
				{
					AddressLine1: "L4, CA House",
					AddressLine2: "14 Boulevard Quay",
					AddressType:  "POBOX",
					City:         "Wellington",
					PostalCode:   "6012",
				},
			},
			Phones: []Phone{
				{
					PhoneType: "DEFAULT",
				},
			},
			IsSupplier: false,
			IsCustomer: false,
		},
		Date:            "\\/Date(1518685950940+0000)\\/",
		DueDate:         "\\/Date(1518685950940+0000)\\/",
		Status:          "AUTHORISED",
		LineAmountTypes: "exclusive",
		LineItems: []LineItem{
			{
				LineItemID:  "52208ff9-528a-4985-a9ad-b2b1d4210e38",
				Description: "Onsite project management ",
				Quantity:    1.00,
				UnitAmount:  1800.00,
				TaxType:     "OUTPUT",
				TaxAmount:   225.00,
				LineAmount:  1800.00,
				AccountCode: "200",
				Tracking: []TrackingCategory{
					{
						TrackingCategoryID: "e2f2f732-e92a-4f3a9c4d-ee4da0182a13",
						Name:               "Activity/Workstream",
					},
				},
			},
		},
		SubTotal:       1800.00,
		TotalTax:       225.00,
		Total:          2025.00,
		UpdatedDateUTC: "\\/Date(1518685950940+0000)\\/",
		CurrencyCode:   "NZD",
		InvoiceID:      "243216c5-369e-4056-ac67-05388f86dc81",
		InvoiceNumber:  "OIT00546",
		Payments: []Payment{
			{
				Date:      "\\/Date(1518685950940+0000)\\/",
				Amount:    1000.00,
				PaymentID: "0d666415-cf77-43fa-80c7-56775591d426",
			},
		},
		AmountDue:      1025.00,
		AmountPaid:     1000.00,
		AmountCredited: 0.00,
	}

	want := `{
	   "Type":"ACCREC",
	   "Contact":{
		  "ContactID":"025867f1-d741-4d6b-b1af-9ac774b59ba7",
		  "ContactStatus":"ACTIVE",
		  "Name":"City Agency",
		  "Addresses":[
			 {
				"AddressType":"STREET"
			 },
			 {
				"AddressType":"POBOX",
				"AddressLine1":"L4, CA House",
				"AddressLine2":"14 Boulevard Quay",
				"City":"Wellington",
				"PostalCode":"6012"
			 }
		  ],
		  "Phones":[
			 {
				"PhoneType":"DEFAULT"
			 }
		  ]
	   },
	   "Date":"\\/Date(1518685950940+0000)\\/",
	   "DueDate":"\\/Date(1518685950940+0000)\\/",
	   "Status":"AUTHORISED",
	   "LineAmountTypes":"exclusive",
	   "LineItems":[
		  {
			 "Description":"Onsite project management ",
			 "Quantity":1,
			 "UnitAmount":1800,
			 "AccountCode":"200",
			 "LineItemID":"52208ff9-528a-4985-a9ad-b2b1d4210e38",
			 "TaxType":"OUTPUT",
			 "TaxAmount":225,
			 "LineAmount":1800,
			 "Tracking":[
				{
				   "TrackingCategoryID":"e2f2f732-e92a-4f3a9c4d-ee4da0182a13",
				   "Name":"Activity/Workstream"
				}
			 ]
		  }
	   ],
	   "SubTotal":1800,
	   "TotalTax":225,
	   "Total":2025,
	   "UpdatedDateUTC":"\\/Date(1518685950940+0000)\\/",
	   "CurrencyCode":"NZD",
	   "InvoiceID":"243216c5-369e-4056-ac67-05388f86dc81",
	   "InvoiceNumber":"OIT00546",
	   "Payments":[
		  {
			 "PaymentID":"0d666415-cf77-43fa-80c7-56775591d426",
			 "Date":"\\/Date(1518685950940+0000)\\/",
			 "Amount":1000
		  }
	   ],
	   "AmountDue":1025,
	   "AmountPaid":1000
	}`

	testJSONMarshal(t, i, want)
}
