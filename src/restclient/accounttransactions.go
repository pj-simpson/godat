package restclient

import (
	"net/http"

	"github.com/pj-simpson/godat/src/models"
)

func (c *Client) GetAccountTransactions(companyId string, connectionId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.AccountTransaction], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetAccountTransactionsURI(companyId, connectionId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.AccountTransaction]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAccountTransaction(companyId string, connectionId string, accountTransactionId string) (*models.AccountTransaction, error) {

	url := c.UrlParser.MakeGetAccountTransactionURI(companyId, connectionId, accountTransactionId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.AccountTransaction{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
