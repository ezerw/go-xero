package xero

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getJson(req *http.Request, target interface{}) error {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error: failed doing the request (%v)", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error: xero responded with %d status code", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(target)
}
