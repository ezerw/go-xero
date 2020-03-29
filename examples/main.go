package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ezerw/go-xero"
)

const (
	AccessToken string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFDQUY4RTY2NzcyRDZEQzAyOEQ2NzI2RkQwMjYxNTgxNTcwRUZDMTkiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJISy1PWm5jdGJjQW8xbkp2MENZVmdWY09fQmsifQ.eyJuYmYiOjE1ODU0NTQzMzMsImV4cCI6MTU4NTQ1NjEzMywiaXNzIjoiaHR0cHM6Ly9pZGVudGl0eS54ZXJvLmNvbSIsImF1ZCI6Imh0dHBzOi8vaWRlbnRpdHkueGVyby5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiRDMzOTIxQTM2RTI1NEZBNTgzQkJCNzJFMkQxMTRBNTgiLCJzdWIiOiJhMmVhOThlNWIxOWI1YjI4OWM1N2ZkZjcxNjdmNmY4ZCIsImF1dGhfdGltZSI6MTU4NTQ1NDMxOSwieGVyb191c2VyaWQiOiI2MzQ1Yzc2MC1kZjk1LTQ0MzYtOGNkMi0xZjZjODZmNzcxNjYiLCJnbG9iYWxfc2Vzc2lvbl9pZCI6IjM3OGRhYTlhOTkyNTQ3YTFhZTI4MDI1MmNlNmJhNjg5IiwianRpIjoiNWE1OWZmNGFhZjk4NWUyNThhNmJlY2RlMTMxMjNkNDEiLCJzY29wZSI6WyJhY2NvdW50aW5nLnNldHRpbmdzIiwiYWNjb3VudGluZy50cmFuc2FjdGlvbnMiLCJhY2NvdW50aW5nLmNvbnRhY3RzIiwib2ZmbGluZV9hY2Nlc3MiXX0.i3aGHh8yDkFGfrdnIziipSvbS8D7K1eC55QGTZcrbtAkPf8eCvAV6shQ7DOqqM2aUVzDdnSXNYIP3ZHuKn4QfZBzX3uvu4rdACsHzUKY03YH5Q6djg78-W6ZIdv_Pd7nNCMUb4HAVc6u3pBYvYGxd4-jbqofN5Ogl72siuosr0XE5KdSs3BfAtPWB2lkWEVEH70D5z-sbZ9ug8lnkZJfnc30F6nEDVq59fQAn0ZqDp3oIM25tHk-J63cLUaASs1a4u8Gjgnt4SxsaBbKHWfEvA-2JO3IyqGpq2MH1xtGmyuzNHdrIlnD63wkpfnzim18bA5PI9GPhzfgtghNuBQ_uA"
	TenantID    string = "0b669781-bbd2-4407-898c-54ecf06ad12b"
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
