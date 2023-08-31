package godat

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pj-simpson/godat/pkg/godat"
	"github.com/pj-simpson/godat/pkg/models"
)

func TestPaginatedResponseParserDefaults(t *testing.T) {

	token := "ARBITRARY"
	fakeURL := "https://api.doesnt.exist/"
	codat := godat.NewCodatClientWithURLOverride(token, fakeURL)
	page := models.PaginatedResponseOptions{}

	pageNum, pageSize, _, _ := codat.PaginatedResponseParser(&page)

	if pageNum != 1 {
		t.Errorf("Expected '1', got %d", pageNum)
	}
	if pageSize != 100 {
		t.Errorf("Expected '100', got %d", pageSize)
	}
}

func TestClientGetCompanies(t *testing.T) {

	//get the companies data
	companiesJSONResp, _ := os.Open("test_data/companies.json")
	companiesJSONRespAsBytes, _ := io.ReadAll(companiesJSONResp)

	// set up the mock server
	mockCodat := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(companiesJSONRespAsBytes)
	}))
	defer mockCodat.Close()

	// set up the Codat client with the mock server's URL
	token := "ARBITRARY"
	fakeURL := mockCodat.URL + "/"
	codat := godat.NewCodatClientWithURLOverride(token, fakeURL)
	page := models.PaginatedResponseOptions{}

	_, err := codat.GetCompanies(&page)
	if err != nil {
		t.Errorf("Get companies mock server failed")
	}

}
