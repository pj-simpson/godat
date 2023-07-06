package godat

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pj-simpson/godat/pkg/models"
	"github.com/pj-simpson/godat/pkg/urlparser"
)

type Client struct {
	APIKey     string
	HTTPClient *http.Client
	UrlParser  *urlparser.UrlParser
}

const (
	CodatProdBaseURL = "https://api.codat.io/"
)

func NewCodatClient(APIKey string) *Client {
	return &Client{
		APIKey:     APIKey,
		HTTPClient: &http.Client{},
		UrlParser: &urlparser.UrlParser{
			BaseURL: CodatProdBaseURL,
		},
	}
}

func (c *Client) PaginatedResponseParser(options *models.PaginatedResponseOptions) (int, int, string, string) {
	pos := options

	var p int
	var ps int
	var q string
	var ob string

	// ensure correct default integers are set

	if pos.Page != 0 {
		p = pos.Page
	} else {
		p = 1
	}

	if pos.PageSize != 0 {
		ps = pos.PageSize
	} else {
		ps = 100
	}

	q = pos.Query
	ob = pos.OrderBy
	return p, ps, q, ob

}

//TODO can we do better than interface here? Generics ?

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+c.APIKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var e models.CodatError
		if err = json.NewDecoder(res.Body).Decode(&e); err == nil {
			return errors.New(e.Service + " " + e.Message + " " + strconv.Itoa(e.StatusCode) + " " + e.CorrelationID)
		}

		return fmt.Errorf("status code from codat system: %d", res.StatusCode)
	}

	response := v

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
