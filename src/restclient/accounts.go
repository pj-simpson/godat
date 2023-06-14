package restclient

import (
	"net/http"

	"github.com/pj-simpson/godat/src/models"
)

func (c *Client) GetAccounts(companyId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.Account], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetAccountsURI(companyId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.Account]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAccount(companyId string, accountId string) (*models.Account, error) {

	url := c.UrlParser.MakeGetAccountURI(companyId, accountId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.Account{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
