package godat

import (
	"net/http"

	"github.com/pj-simpson/godat/src/models"
)

func (c *Client) GetBankAccounts(companyId string, connectionId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.BankAccount], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetBankAccountsURI(companyId, connectionId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.BankAccount]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetBankAccount(companyId string, connectionId string, bankAccountId string) (*models.BankAccount, error) {

	url := c.UrlParser.MakeGetBankAccountURI(companyId, connectionId, bankAccountId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.BankAccount{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
