package models

type AccountTransaction struct {
	ID                 string                   `json:"id"`
	TransactionID      string                   `json:"transactionId"`
	Note               string                   `json:"note"`
	BankAccountRef     BankAccountRef           `json:"bankAccountRef"`
	Date               string                   `json:"date"`
	Status             string                   `json:"status"`
	Currency           string                   `json:"currency"`
	CurrencyRate       int                      `json:"currencyRate"`
	Lines              []AccountTransactionLine `json:"lines"`
	TotalAmount        float64                  `json:"totalAmount"`
	ModifiedDate       string                   `json:"modifiedDate"`
	SourceModifiedDate string                   `json:"sourceModifiedDate"`
	Metadata           Metadata                 `json:"metadata"`
}

type AccountTransactionLine struct {
	Description string    `json:"description"`
	RecordRef   RecordRef `json:"recordRef"`
	Amount      float64   `json:"amount"`
}

type BankAccountRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
