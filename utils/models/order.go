package models

type OrderDetails struct {
	OrderId        string
	FinalPrice     float64
	ShipmentStatus string
	PaymentStatus  string
}

type OrderProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}
type FullOrderDetails struct {
	OrderDetails        OrderDetails
	OrderProductDetails []OrderProductDetails
}
type OrderProducts struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type Invoice struct {
	Cart        []Cart
	AddressInfo AddressInfoResponse
}

// ORDER DETAILS

type CombinedOrderDetails struct {
	OrderId        string  `json:"order_id"`
	FinalPrice     float64 `json:"final_price"`
	ShipmentStatus string  `json:"shipment_status"`
	PaymentStatus  string  `json:"payment_status"`
	Name           string  `json:"firstname"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	HouseName      string  `json:"house_name" validate:"required"`
	State          string  `json:"state" validate:"required"`
	Pin            string  `json:"pin" validate:"required"`
	Street         string  `json:"street"`
	City           string  `json:"city"`
}
