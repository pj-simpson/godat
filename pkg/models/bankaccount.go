package models

type BankAccount struct {
	ID                 string   `json:"id"`
	AccountName        string   `json:"accountName"`
	AccountType        string   `json:"accountType"`
	SortCode           string   `json:"sortCode"`
	AccountNumber      string   `json:"accountNumber"`
	Currency           string   `json:"currency"`
	Balance            float64  `json:"balance"`
	ModifiedDate       string   `json:"modifiedDate"`
	SourceModifiedDate string   `json:"sourceModifiedDate"`
	Metadata           Metadata `json:"metadata"`
}
