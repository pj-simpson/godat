package main

import (
	"fmt"
	"os"

	"github.com/pj-simpson/godat/src/models"
	"github.com/pj-simpson/godat/src/restclient"
)

///companies/f0c043aa-498e-4f4f-b5fc-872951a36550/data/accounts/2

func main() {
	token := os.Getenv("CODAT_TOKEN")
	codat := restclient.NewCodatClient(token)
	page := &models.PaginatedResponseOptions{
		Page:     1,
		PageSize: 25,
	}
	// company := "9e04feaf-10d8-49fc-a9ed-05f5870e8405"
	// connection := "abd9cfbe-22e4-4a4e-90af-2e8c1cdbff97"
	// acc := "ba472be2-fdfe-433b-9acc-5a7f2a37dd56"

	trans, err := codat.GetCompanies(page)

	if err != nil {
		fmt.Printf("%#v", err)
	}

	fmt.Printf("%#v", trans)
}
