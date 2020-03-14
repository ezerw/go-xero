package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ezerw/go-xero"
)

const (
	AccessToken string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODQxNTkyMjIsImV4cCI6MTU4NDE2MTAyMiwiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4NDE1OTIxNCwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6IjBiNDIyYzcyNzhlODQxNDliZTc0YzdkOTFlNDM1YmVhIiwianRpIjoiMGUyNzZmOTkxMzIzZDE1YjBjYzViN2QyYmFkMGQ2YzciLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.rChczHqggVJDNha__YmmY5rAtZP07r5wHukh05jUh52tHjU3DjLlsfSXF70Dz7c7e13HDLYhmgZwFq1CIC3MZ8PrmgaNtvLaJlcfuVtPBC1hpS-cufGcDhFVQIoedjdltY4pTlStn6ApfaWyndE_jQTwbQcFWt247NyZVOqcUXvxCmIPugegD2pMFOKkE4tu3Uaqnv2LawrvmGKA8rvEWbZj7Qjlm1QZS4M_A5NsaXSyHbh0oOdOJr2XIibMNoMQ6zlN_h4Etx4qOqZFHphn0DLDpWYPtaVR9uYa7OKl7iQJx8iikKUQNYw2UI-rMHX2zX4ozyQUFGZe76_xQEtdhg"
	TenantID    string = "f0c0ee1f-4091-41f9-8f02-0309a955da3d"
)

func main() {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: AccessToken})
	client := xero.NewClient(
		oauth2.NewClient(ctx, tokenSource),
		xero.TenantID(TenantID),
	)

	// Create an Account
	account, err := createAccount(ctx, client)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", account)
}
