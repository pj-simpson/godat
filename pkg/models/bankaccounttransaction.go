package models

type BankAccountTransaction struct {
	AccountID       string  `json:"accountId"`
	ClearedOnDate   string  `json:"clearedOnDate"`
	Description     string  `json:"description"`
	Counterparty    string  `json:"counterparty"`
	Reconciled      bool    `json:"reconciled"`
	Amount          float64 `json:"amount"`
	Balance         float64 `json:"balance"`
	TransactionType string  `json:"transactionType"`
	ModifiedDate    string  `json:"modifiedDate"`
}
