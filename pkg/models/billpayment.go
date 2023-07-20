package models

type BillPaymentLine struct {
	Amount          float64           `json:"amount"`
	Links           []PaymentLineLink `json:"links"`
	AllocatedOnDate string            `json:"allocatedOnDate"`
}

type PaymentLineLink struct {
	Type         string  `json:"type"`
	ID           string  `json:"id"`
	Amount       float64 `json:"amount"`
	CurrencyRate float64 `json:"currencyRate"`
}

type BillPayment struct {
	ID                 string            `json:"id"`
	SupplierRef        SupplierRef       `json:"supplierRef"`
	AccountRef         AccountRef        `json:"accountRef"`
	TotalAmount        float64           `json:"totalAmount"`
	Currency           string            `json:"currency"`
	CurrencyRate       float64           `json:"currencyRate"`
	Date               string            `json:"date"`
	Note               string            `json:"note"`
	PaymentMethodRef   PaymentMethodRef  `json:"paymentMethodRef"`
	Lines              []BillPaymentLine `json:"lines"`
	ModifiedDate       string            `json:"modifiedDate"`
	SourceModifiedDate string            `json:"sourceModifiedDate"`
	Reference          string            `json:"reference"`
}
