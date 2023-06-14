package models

type PaginatedResponseOptions struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Query    string `json:"query"`
	OrderBy  string `json:"orderBy"`
}

type PaginatedResponse[C any] struct {
	Results      []C   `json:"results"`
	PageNumber   int   `json:"pageNumber"`
	PageSize     int   `json:"pageSize"`
	TotalResults int   `json:"totalResults"`
	Links        Links `json:"_links"`
}

type Link struct {
	Href string `json:"href"`
}

type Links struct {
	Current Link `json:"current"`
	Self    Link `json:"self"`
	Next    Link `json:"next"`
}

type CodatError struct {
	StatusCode        int    `json:"statusCode"`
	Service           string `json:"service"`
	Message           string `json:"error"`
	CorrelationID     string `json:"correlationId"`
	CanBeRetried      string `json:"canBeRetried"`
	DetailedErrorCode int    `json:"detailedErrorCode"`
}

type ValidDatatypeLinks struct {
	Property string   `json:"property"`
	Links    []string `json:"links"`
}

type Metadata struct {
	IsDeleted bool `json:"isDeleted"`
}

type RecordRef struct {
	ID       string `json:"id"`
	DataType string `json:"dataType"`
}

type SupplierRef struct {
	ID           string `json:"id"`
	SupplierName string `json:"supplierName"`
}

type AccountRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TaxRateRef struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	EffectiveTaxRate int    `json:"effectiveTaxRate"`
}

type ItemRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrackingCategoryRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CustomerRef struct {
	ID          string `json:"id"`
	CompanyName string `json:"companyName"`
}

type ProjectRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WithholdingTax struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Payment struct {
	ID           string     `json:"id"`
	Note         string     `json:"note"`
	Reference    string     `json:"reference"`
	AccountRef   AccountRef `json:"accountRef"`
	Currency     string     `json:"currency"`
	CurrencyRate int        `json:"currencyRate"`
	PaidOnDate   string     `json:"paidOnDate"`
	TotalAmount  int        `json:"totalAmount"`
}

type Allocation struct {
	Currency        string `json:"currency"`
	CurrencyRate    int    `json:"currencyRate"`
	AllocatedOnDate string `json:"allocatedOnDate"`
	TotalAmount     int    `json:"totalAmount"`
}

type PaymentAllocation struct {
	Payment    Payment    `json:"payment"`
	Allocation Allocation `json:"allocation"`
}
