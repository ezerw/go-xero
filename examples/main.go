package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"

	"github.com/ezerw/go-xero"
)

const (
	AccessToken string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODQxNDcxMDMsImV4cCI6MTU4NDE0ODkwMywiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4NDE0NzA5NCwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6IjgyY2IwNDY1ZmVmMjRhYTRiZmM4MGNmZjI4OTAyN2NiIiwianRpIjoiNzVhMjQ3YmJhZWY4ZDA1YWNlNjBjZjRhOWYxZmQ3MDUiLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.mVGP_RKEmQOyjYxQrhOFE1AHMcNaEmGy2GtSXUiHTK4knhuFed98O6qZWUKQTLGPdw7KUWI9vqcjGJSNGvMtwkYUCYcAzMjRkmr0HVkgpMmMWewGKtd_Grlg67qP8z_Cb5ZYDTBpMAhDaYXdZ_jcXD9eVy3UIqwCwBwqwGV2U9c2s4BAwE9lzK5Nl6XTbxlAZ2QQxszPpLKDVKPUxBJjFQ6kbmxU6v25K79oKSWzvza4MRdG7S8UHDjwcVVG_wEgSC1hZtk65-0auxCdXTqkaPiQutH0jtZWNHtg15wdp9RmJ7pKTCKYtOvvlJbEmCAdO_qvQxVISDOs2bWCjLaHKA"
	TenantID    string = "f0c0ee1f-4091-41f9-8f02-0309a955da3d"
)

func main() {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: AccessToken})
	client := xero.NewClient(
		oauth2.NewClient(ctx, tokenSource),
		xero.TenantID(TenantID),
	)

	//Create an invoice
	i, err := create(ctx, client)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Invoice created. #%s", i.InvoiceNumber)
}
