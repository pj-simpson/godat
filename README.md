# GoDat #

GoDat is an Unofficial Go library for interacting with the [Codat API](https://docs.codat.io/).

## Installation ##

```bash
go get github.com/pj-simpson/godat
```

## Usage ##

Get the Codat token from an enviroment variable.
Create a new Codat REST client. 
Set the pagination query params via the PaginatedResponseOptions struct, to obtain 25 comapanies. 
Call the 'Get Companies' method and print out the result. 

```go
package main

import (
	"fmt"
	"os"

	"github.com/pj-simpson/godat/pkg/godat"
	"github.com/pj-simpson/godat/pkg/models"
)

func main() {
	token := os.Getenv("CODAT_TOKEN")
	codat := godat.NewCodatClient(token)
	page := models.PaginatedResponseOptions{
		Page:     1,
		PageSize: 25,
	}

	comps, err := codat.GetCompanies(&page)

	if err != nil {
		fmt.Printf("%#v", err)
	}

	fmt.Printf("%#v", comps)
}


```
## License ##

MIT

## Caveats ##

Medium term aim for this package is to wrap the entire [Codat Accounting API](https://docs.codat.io/accounting-api#/). 