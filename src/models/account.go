package models

type Account struct {
	ID                     string               `json:"id"`
	NominalCode            string               `json:"nominalCode"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	FullyQualifiedCategory string               `json:"fullyQualifiedCategory"`
	FullyQualifiedName     string               `json:"fullyQualifiedName"`
	Currency               string               `json:"currency"`
	CurrentBalance         float64              `json:"currentBalance"`
	Type                   string               `json:"type"`
	Status                 string               `json:"status"`
	IsBankAccount          bool                 `json:"isBankAccount"`
	ModifiedDate           string               `json:"modifiedDate"`
	SourceModifiedDate     string               `json:"sourceModifiedDate"`
	ValidDatatypeLinks     []ValidDatatypeLinks `json:"validDatatypeLinks"`
}
