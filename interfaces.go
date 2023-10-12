package main


type HelloWorld struct {
	Message string `json:"message"`
}

type OrderDetail struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	VariantName string  `json:"variantName"`
	Qty         float64 `json:"qty"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Pot         float64 `json:"pot"`
	Subtotal    float64 `json:"subtotal"`
}

type CustomerOrder struct {
	ID            uint64        `json:"id"`
	SalesName     string        `json:"salesName,omitempty"`
	CreatedAt     string        `json:"createdAt"`
	DueAt         string        `json:"dueAt"`
	Description   string        `json:"description,omitempty"`
	Total         float64       `json:"total"`
	Payment       float64       `json:"payment"`
	RemainPayment float64       `json:"remainPayment"`
	CustomerName  string        `json:"customerName"`
	Address       string        `json:"address"`
	UpdatedBy     string        `json:"updatedBy"`
	Details       []OrderDetail `json:"details"`
}
