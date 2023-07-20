package godat

import (
	"net/http"

	"github.com/pj-simpson/godat/pkg/models"
)

func (c *Client) GetBillCreditNotes(companyId string, options *models.PaginatedResponseOptions) (*models.PaginatedResponse[models.BillCreditNote], error) {

	page, pageSize, query, orderBy := c.PaginatedResponseParser(options)
	url := c.UrlParser.MakeGetBillCreditNotesURI(companyId, page, pageSize, query, orderBy)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.PaginatedResponse[models.BillCreditNote]{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetBillCreditNote(companyId string, billCreditNoteId string) (*models.BillCreditNote, error) {

	url := c.UrlParser.MakeGetBillCreditNoteURI(companyId, billCreditNoteId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res := models.BillCreditNote{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
