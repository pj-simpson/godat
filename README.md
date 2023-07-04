# GoDat #

GoDat is an Unofficial Go library for interacting with the [Codat API](https://docs.codat.io/accounting-api#/).

## Installation ##

```bash
go get github.com/pj-simpson/godat
```

## Usage ##

Get the Codat token from an enviroment variable.
Create a new Codat REST client. 
Set the pagination query params via the PaginatedResponseOptions struct. 
Call the 'Get Companies' method. 
This will return a single page of 25 companies from the Codat system, 
parsed into a go type. 

```go
package main

import (
	"fmt"
	"os"

	"github.com/pj-simpson/godat"
)

func main() {
	token := os.Getenv("CODAT_TOKEN")
	codat := godat.NewCodatClient(token)
	page := godat.PaginatedResponseOptions{
		Page:     1,
		PageSize: 25,
	}

	comps, err := codat.GetCompanies(page)

	if err != nil {
		fmt.Printf("%#v", err)
	}

	fmt.Printf("%#v", comps)
}

```
## License ##

MIT

## Caveats ##

This package is a POC really, (currently v 0.0.1), my medium term aim is to wrap the entire Accounting API.  