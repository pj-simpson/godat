package restclient

import (
	"net/http"

	"github.com/pj-simpson/godat/src/models"
)

func (c *Client) GetBankAccountTransactions(companyId string, connectionId string, accountId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.BankAccountTransaction], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetBankAccountTransactionsURI(companyId, connectionId, accountId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.BankAccountTransaction]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
