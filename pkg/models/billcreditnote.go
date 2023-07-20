package models

type BillCreditNote struct {
	ID                   string                   `json:"id"`
	BillCreditNoteNumber string                   `json:"billCreditNoteNumber"`
	SupplierRef          SupplierRef              `json:"supplierRef"`
	WithholdingTax       []WithholdingTax         `json:"withholdingTax"`
	TotalAmount          float64                  `json:"totalAmount"`
	TotalDiscount        float64                  `json:"totalDiscount"`
	SubTotal             float64                  `json:"subTotal"`
	TotalTaxAmount       float64                  `json:"totalTaxAmount"`
	DiscountPercentage   float64                  `json:"discountPercentage"`
	RemainingCredit      float64                  `json:"remainingCredit"`
	Status               string                   `json:"status"`
	IssueDate            string                   `json:"issueDate"`
	AllocatedOnDate      string                   `json:"allocatedOnDate"`
	Currency             string                   `json:"currency"`
	CurrencyRate         float64                  `json:"currencyRate"`
	LineItems            []BillCreditNoteLineItem `json:"lineItems"`
	PaymentAllocations   []PaymentAllocation      `json:"paymentAllocations"`
	ModifiedDate         string                   `json:"modifiedDate"`
	SourceModifiedDate   string                   `json:"sourceModifiedDate"`
	Note                 string                   `json:"note"`
}

type BillCreditNoteLineItem struct {
	Description          string                `json:"description"`
	UnitAmount           float64               `json:"unitAmount"`
	Quantity             float64               `json:"quantity"`
	DiscountAmount       float64               `json:"discountAmount"`
	SubTotal             float64               `json:"subTotal"`
	TaxAmount            float64               `json:"taxAmount"`
	TotalAmount          float64               `json:"totalAmount"`
	DiscountPercentage   float64               `json:"discountPercentage"`
	AccountRef           AccountRef            `json:"accountRef"`
	TaxRateRef           TaxRateRef            `json:"taxRateRef"`
	ItemRef              ItemRef               `json:"itemRef"`
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs"`
	Tracking             BillTracking          `json:"tracking"`
	IsDirectCost         bool                  `json:"isDirectCost"`
}
