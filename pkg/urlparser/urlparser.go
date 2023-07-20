package urlparser

import (
	"fmt"
)

type UrlParser struct {
	BaseURL string
}

// Companies

func (u *UrlParser) MakeGetCompaniesURI(page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, page, pageSize, query, orderBy)
}

// Accounts
func (u *UrlParser) MakeGetAccountsURI(companyId string, page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies/%s/data/accounts?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, companyId, page, pageSize, query, orderBy)
}

func (u *UrlParser) MakeGetAccountURI(companyId string, accountId string) string {
	return fmt.Sprintf("%scompanies/%s/data/accounts/%s", u.BaseURL, companyId, accountId)
}

// Account Transactions

func (u *UrlParser) MakeGetAccountTransactionsURI(companyId string, connectionId string, page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies/%s/connections/%s/data/accountTransactions?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, companyId, connectionId, page, pageSize, query, orderBy)
}

func (u *UrlParser) MakeGetAccountTransactionURI(companyId string, connectionId string, accountTransactionId string) string {
	return fmt.Sprintf("%scompanies/%s/connections/%s/data/accountTransactions/%s", u.BaseURL, companyId, connectionId, accountTransactionId)
}

// Bank Accounts

func (u *UrlParser) MakeGetBankAccountsURI(companyId string, connectionId string, page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies/%s/connections/%s/data/bankAccounts?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, companyId, connectionId, page, pageSize, query, orderBy)
}

func (u *UrlParser) MakeGetBankAccountURI(companyId string, connectionId string, bankAccountId string) string {
	return fmt.Sprintf("%scompanies/%s/connections/%s/data/bankAccounts/%s", u.BaseURL, companyId, connectionId, bankAccountId)
}

// Bank Account Transactions
func (u *UrlParser) MakeGetBankAccountTransactionsURI(companyId string, connectionId string, accountId string, page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies/%s/connections/%s/data/bankAccounts/%s/bankTransactions?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, companyId, connectionId, accountId, page, pageSize, query, orderBy)
}

// Bills

func (u *UrlParser) MakeGetBillsURI(companyId string, page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies/%s/data/bills?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, companyId, page, pageSize, query, orderBy)
}

func (u *UrlParser) MakeGetBillURI(companyId string, billId string) string {
	return fmt.Sprintf("%scompanies/%s/data/bills/%s", u.BaseURL, companyId, billId)
}

// Bill Credit Notes

func (u *UrlParser) MakeGetBillCreditNotesURI(companyId string, page int, pageSize int, query string, orderBy string) string {
	return fmt.Sprintf("%scompanies/%s/data/billCreditNotes?page=%d&pageSize=%d&query=%s&orderby=%s", u.BaseURL, companyId, page, pageSize, query, orderBy)
}

func (u *UrlParser) MakeGetBillCreditNoteURI(companyId string, billCreditNoteId string) string {
	return fmt.Sprintf("%scompanies/%s/data/billCreditNotes/%s", u.BaseURL, companyId, billCreditNoteId)
}
