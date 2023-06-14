package restclient

import (
	"net/http"

	"github.com/pj-simpson/godat/src/models"
)

func (c *Client) GetCompanies(options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.Company], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetCompaniesURI(page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.Company]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
