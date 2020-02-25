package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ezerw/xero"
)

const (
	AccessToken string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODI1OTExMDcsImV4cCI6MTU4MjU5MjkwNywiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4MjU5MTA5OCwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6ImZhMjdhMDMwYmM0OTRkZmZhNWQzZDMxOTA0MjYxOTI2IiwianRpIjoiNDMyNGQzYzVmNjk4MTM5MTU2ODc4MDk2MjIzZDI0MjgiLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.JBpeXa379ZYNU1yLXH9VYqGAU-AgbhhFoT3qjZvHynst7ZjdkKBGlviDuzJW098R5RThizbaczY9GESmtMFwMUjGl2Lt-IHG5A1hWY0JXUKODD0PVUhD-E3CoNET9uGUVpbmmZDNIW9FxKoBmnCQzA3_8swvgP90xz4-Scqdy6tJ8v3kBMerWgdcRj6_neU2z5EWE8eS7ojcmVSP1HbUT6oguPbRdYdl-cqsBSqcWuymPs4aqoAGROvAMEDTHTiv7bbtrbmdGlmCno-GAwYhEFSD8Z5gvTFX9INwdnWWkiMCuPXZ3XCN4Jvdd9qe97EyGIR6uD4WBr984HA6uzhCYQ"
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
