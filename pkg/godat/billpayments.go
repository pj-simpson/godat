package godat

import (
	"net/http"

	"github.com/pj-simpson/godat/pkg/models"
)

func (c *Client) GetBillPayments(companyId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.BillPayment], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetBillPaymentsURI(companyId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.BillPayment]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetBillPayment(companyId string, billPaymentId string) (*models.BillPayment, error) {

	url := c.UrlParser.MakeGetBillPaymentURI(companyId, billPaymentId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.BillPayment{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
