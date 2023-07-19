package godat

import (
	"net/http"

	"github.com/pj-simpson/godat/pkg/models"
)

func (c *Client) GetBills(companyId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.Bill], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetBillsURI(companyId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.Bill]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetBill(companyId string, billId string) (*models.Bill, error) {

	url := c.UrlParser.MakeGetBillURI(companyId, billId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.Bill{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
