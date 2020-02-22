package xero

import (
	"fmt"
	"time"
)

const ApiUrl = "https://api.xero.com/api.xro/2.0"

type (
	// Config types
	AccessToken string
	TenantID    string

	// Config structure
	Config struct {
		AccessToken AccessToken
		TenantID    TenantID
		Timeout     time.Duration
	}

	// General client
	client struct {
		Config *Config
	}
)

func NewClient(config *Config) (*client, error) {
	if config.AccessToken == "" || config.TenantID == "" {
		return nil, fmt.Errorf("error: required AccessToken and TenantID were not provided")
	}

	return &client{Config: config}, nil
}
