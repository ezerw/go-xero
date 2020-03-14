# go-xero

go-xero is a Go client library for accessing the [Xero API](https://developer.xero.com/documentation/oauth2/overview).

### Authentication ###
The go-xero library does not directly handle authentication. 
Instead, when creating a new client, pass an `http.Client` that can handle authentication for you. 
The easiest and recommended way to do this is using the [oauth2](https://github.com/golang/oauth2) library, 
but you can always use any other library that provides an `http.Client`.

Note that a `XERO_TENANT_ID` is required creating the Xero client instance.
If you have an OAuth2 `XERO_ACCESS_TOKEN`, you can use it with the oauth2 library using:

```go
import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/ezerw/go-xero"
)

func main() {
    tokenSource := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "ACCESS_TOKEN"}
    )
    
    ctx := context.Background()
    httpClient := oauth2.NewClient(ctx, tokenSource)
    
    xeroClient := xero.NewClient(httpClient, "XERO_TENANT_ID")
}
````

#### Please refer to Wiki for examples:
[Invoices](https://github.com/ezerw/go-xero/wiki/Invoices)