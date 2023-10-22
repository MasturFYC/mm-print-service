package main

type DeliveryOrder struct {
	Delivery_id uint64 `json:"delivery_id"`
	Created_at  string `json:"created_at"`
	Description string `json:"description"`
	Driver_id   uint64 `json:"driver_id"`
	Nopol       string `json:"nopol"`
	User_name   string `json:"user_name"`
	Is_shipped  bool   `json:"is_shipped"`
}

// type DeliveryWithOrder struct {
// 	Order_id   uint64  `json:"order_id"`
// 	Created_at string  `json:"created_at"`
// 	Payment_id string  `json:"payment_id"`
// 	Updated_by string  `json:"updated_by"`
// 	Due_at     string  `json:"due_at"`
// 	Total      float64 `json:"total"`
// 	Sales      string  `json:"sales"`
// 	Customer   string  `json:"customer"`
// }

type DeliveryWithProduct struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	Variant_name string  `json:"variant_name"`
	Qty          float64 `json:"qty"`
	Unit         string  `json:"unit"`
}

type DeliveryOrderList struct {
	DeliveryOrder
	Driver_name string `json:"driver_name"`
}
type PrintDelivery struct {
	Delivery DeliveryOrderList     `json:"delivery"`
	Details  []DeliveryWithProduct `json:"details"`
}

type PrintDataPayment struct {
	PaymentId    uint64  `json:"paymentId"`
	OrderId      uint64  `json:"orderId"`
	CreatedAt    string  `json:"createdAt"`
	Amount       float64 `json:"amount"`
	Remain       float64 `json:"remain"`
	CustomerName string  `json:"customerName"`
	Description  string  `json:"description"`
	Terbilang    string  `json:"terbilang"`
	Admin        string  `json:"admin"`
}
