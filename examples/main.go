package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ezerw/xero"
)

const (
	AccessToken string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODI1MDk5NTQsImV4cCI6MTU4MjUxMTc1NCwiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4MjUwOTk0NSwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6ImViYzEzMGE0OWI3NzQ3Y2RiMzJlYmMxZDMzZjdkNDAzIiwianRpIjoiZmVmY2Q1ZjViN2NkMmNjZmVjOGMwZjU4MDg2ZjdlM2UiLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.SjMF_E6ekO9we3yryDH_TiwmxWj8L0VjiOGMcju9_6msjPtxc3gwGOZlZTRHlm21-D6mCR658tJDEY4-39o1QMYc7jsqPGvfYSUg3Qd8K_pkYGuJrU9Bce8d74U46IvYmoWJcGhW1gWaHrDyH7eoGpfT3D-igXggQqBwWwgiV0XJSh0cMEjhXe0crMCi3WBiCSKOZyHFjb69y60i4VQe2YIIpTpZZuiydsj3-jT25jJ_O05UdGH5m4bkOMZQQJJGXUgHPc7G3u_XJ-Xm3AXVa6zI1QxOdgbVEgqCBM75xXsgNnixCAW3DvFjf5O2JEX73LEnTm-rS_QLQTP3Q5yu8A"
	TenantID    string = "f0c0ee1f-4091-41f9-8f02-0309a955da3d"
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
