package xero

import (
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.xero.com/api.xro/2.0"
)

// TenantID used when communicating with the Xero API.
type TenantID string

type service struct {
	client *Client
}

// A Client manages communication with the Xero API.
type Client struct {
	client   *http.Client

	// Base URL for API requests.
	BaseURL  *url.URL

	// Tenant used to do requests to API endpoints.
	TenantID TenantID

	// Services used for talking to different parts of the Xero API.
	Invoices *InvoicesService

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service
}

// NewClient returns a new Xero API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(httpClient *http.Client, tenantID TenantID) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, TenantID: tenantID}
	c.common.client = c

	// Available Xero API resources
	c.Invoices = (*InvoicesService)(&c.common)

	return c
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	//TODO: handle rate limit, custom response errors, etc
	return c.client.Do(req)
}
