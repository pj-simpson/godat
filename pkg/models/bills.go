package models

type PurchaseOrderRefs struct {
	ID                  string `json:"id"`
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
}

type InvoiceTo struct {
	ID       string `json:"id"`
	DataType string `json:"dataType"`
}

type BillTracking struct {
	RecordRefs   []RecordRef `json:"recordRefs"`
	CustomerRef  CustomerRef `json:"customerRef"`
	InvoiceTo    InvoiceTo   `json:"invoiceTo"`
	ProjectRef   ProjectRef  `json:"projectRef"`
	IsBilledTo   string      `json:"isBilledTo"`
	IsRebilledTo string      `json:"isRebilledTo"`
}

type BillLineItems struct {
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

type Bill struct {
	ID                 string              `json:"id"`
	Reference          string              `json:"reference"`
	SupplierRef        SupplierRef         `json:"supplierRef"`
	PurchaseOrderRefs  []PurchaseOrderRefs `json:"purchaseOrderRefs"`
	IssueDate          string              `json:"issueDate"`
	DueDate            string              `json:"dueDate"`
	Currency           string              `json:"currency"`
	CurrencyRate       float64             `json:"currencyRate"`
	LineItems          []BillLineItems     `json:"lineItems"`
	WithholdingTax     []WithholdingTax    `json:"withholdingTax"`
	Status             string              `json:"status"`
	SubTotal           float64             `json:"subTotal"`
	TaxAmount          float64             `json:"taxAmount"`
	TotalAmount        float64             `json:"totalAmount"`
	AmountDue          float64             `json:"amountDue"`
	ModifiedDate       string              `json:"modifiedDate"`
	SourceModifiedDate string              `json:"sourceModifiedDate"`
	Note               string              `json:"note"`
	PaymentAllocations []PaymentAllocation `json:"paymentAllocations"`
	Metadata           Metadata            `json:"metadata"`
}
