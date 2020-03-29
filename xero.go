package xero

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	baseURL = "https://api.xero.com/api.xro/2.0"

	// Xero Rate limit response headers
	// https://developer.xero.com/documentation/oauth2/limits
	headerRateDayLimitRemain = "X-DayLimit-Remaining" // Number of remaining Day limit.
	headerRateMinLimitRemain = "X-MinLimit-Remaining" // Number of remaining Minute limit.
	headerRateLimit          = "X-Rate-Limit-Problem" // Which limit you have reached.
	headerRateRetry          = "Retry-After"          // How many seconds to wait before making another request.

	mediaTypeJSON = "application/json"
)

// TenantID used when communicating with the Xero API.
type TenantID string

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// A Client manages communication with the Xero API.
type Client struct {
	client HTTPClient

	// Base URL for API requests.
	BaseURL *url.URL

	// Tenant used to do requests to API endpoints.
	TenantID TenantID

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Services used for talking to different parts of the Xero API.
	Invoices *InvoicesService
	Accounts *AccountsService
}

type service struct {
	client *Client
}

// NewClient returns a new Xero API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(httpClient HTTPClient, tenantID TenantID) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, TenantID: tenantID}
	c.common.client = c

	// Available Xero API resources
	c.Invoices = (*InvoicesService)(&c.common)
	c.Accounts = (*AccountsService)(&c.common)

	return c
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden:
		invalidTokenError := &InvalidTokenError{}
		err := parseResponse(res, invalidTokenError)
		if err != nil {
			return nil, err
		}
		return nil, invalidTokenError
	case http.StatusBadRequest:
		badRequestError := &BadRequestError{}
		err := parseResponse(res, badRequestError)
		if err != nil {
			return nil, err
		}

		// Return a JSON string with the list of errors for more detailed view of what happened.
		b, err := json.Marshal(badRequestError)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}

	// For POST and PUT requests
	if v != nil {
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			if !errors.Is(err, io.EOF) { // ignore EOF errors caused by empty response body
				return nil, err
			}
		}
	}

	return res, nil
}

// NewRequest creates an API request. If specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method string, url string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", mediaTypeJSON)
	req.Header.Set("Xero-tenant-id", string(c.TenantID))

	if body != nil {
		req.Header.Set("Content-Type", mediaTypeJSON)
	}

	return req, nil
}

// parseResponse decodes the response body into the target
func parseResponse(res *http.Response, target interface{}) error {
	err := json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("failed to read error response body: %s", err)
	}
	return nil
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
