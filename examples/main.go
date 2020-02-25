package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ezerw/xero"
)

const (
	AccessToken string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODI2MDE3NjksImV4cCI6MTU4MjYwMzU2OSwiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4MjYwMTc2MCwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6ImFhZGExZDlhMTRlMTQ5ODc4MTY3NWQ4NjdkNDZmYTZiIiwianRpIjoiYjI3ODYwMTRkYjE0ZmI1YmYyNWM0N2JkYWY4NTE1YjEiLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.AIuvOUW7BMc1jIdL1YGW4r9RNleZLWkjEg65uyoMUNvRMDq-1-9eYkmFutVpviSQMxVpE2bWvd6hbHr2WzMOfYzZXa5lP0i1i9Abs8YyJq1RshQK83B2miHQCJhr9a9RXPSJeAPqQDrKa3S94dYX9V7opqAM2ld1IpnUE9q_IEfdBXoyLw3478mvcXMRhvU-XxN6-zHnaXWjwIh8DdCrvKqxdcskuCLlKdf34V1pM0lmqaggnDlbaHVDF6Z_HZMf2eH-k6ugcbWuFxy7kp8R7mxNdDwud1fMpazqwIWXhZexGcUlXa5UD_3tKQYU30_IWkB-ZeRcpiCcitFGWPIviA"
	TenantID    string = "f0c0ee1f-4091-41f9-8f02-0309a955da3d"
)

func main() {
	ctx := context.Background()

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: AccessToken})
	httpClient := oauth2.NewClient(ctx, tokenSource)

	xeroClient := xero.NewClient(httpClient, xero.TenantID(TenantID))

	//invoices, err := xeroClient.Invoices.List(ctx)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("error: %v", err))
	//	return
	//}

	invoice, err := xeroClient.Invoices.GetByID(ctx, "ac1fc284-2e69-46b6-95b3-17949c65bb42")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}

	fmt.Printf("%+v\n", invoice)
}
