package main

import (
	"fmt"
	"time"

	"github.com/ezerw/xero"
)

func main() {
	config := &xero.Config{
		AccessToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODIzNTc5MjYsImV4cCI6MTU4MjM1OTcyNiwiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4MjM1NzkxNiwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6ImIzZTMzMmExOGI2ODRjMzZiNjdiNmViNDZmMmQ0YWQzIiwianRpIjoiNzVmZDVkYTBlNGFhNGRkZGU3NGFlYmI1NzI5ODMwNDQiLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.JORWO73Td_7hKdZvUfwiP6hXfXDkXrjhvUZoEVtm6xnalTdKYuJrHZ9gsIFfAdrJKMj4Tu4g624M2XpU7u4IxNduK61vWBwnSCEWLCKnbKnHMNt0O2tpHTa4IHjTHDu7OxrRmHbX81-vINns6HmjV0CnHqOOlPD3y5arM1t3qFMW5cn49eIwow-ulnWTd8rsB9JZ8uRh69CQFVHKnzXkFfbrw8wzV0FbpUlBwcxzMVhBIApdS0Qxj_ulYbpIvRRpff3ferFh4_h_Ti05U0zx5zaei2NrI42-SL4y44vex2IA9Wd-7zZ6Oh9FdcVPXDflm2R5TVtiSDABSRO_MEuFJw",
		TenantID:    "f0c0ee1f-4091-41f9-8f02-0309a955da3d",
		Timeout:     5 * time.Second, // Request timeout
	}

	xeroAPI, err := xero.NewClient(config)
	if err != nil {
		fmt.Println(fmt.Errorf("error creating the client: %v", err))
	}

	// Find one invoice
	invoice, err := xeroAPI.Invoices().Find("22b3fab4-ef56-4d70-a110-a7cc3c1a26cd")
	if err != nil {
		fmt.Println(fmt.Errorf("error getting invoice: %v", err))
	}

	newInvoice := xero.Invoice{
		Reference: "Eze",
	}

	uInvoice, err := xeroAPI.Invoices().Update(invoice, newInvoice)
	fmt.Printf("%+v\n", uInvoice)
}
