package domain

import "time"

type Order struct {
	OrderId         string        `json:"order_id" gorm:"primaryKey;not null"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	DeletedAt       *time.Time    `json:"deleted_at" gorm:"index"`
	DeliveryTime    time.Time     `json:"delivery_time"`
	UserID          int           `json:"user_id" gorm:"not null"`
	User            User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	GrandTotal      float64       `json:"grand_total"`
	FinalPrice      float64       `json:"discount_price"`
	ShipmentStatus  string        `json:"shipment_status"`
	PaymentStatus   string        `json:"payment_status"`
	Approval        bool          `json:"approval"`
}

type OrderItem struct {
	ID         uint     `json:"id" gorm:"primaryKey;not null"`
	OrderID    string   `json:"order_id"`
	Order      Order    `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID  uint     `json:"product_id"`
	Products   Products `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   int      `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}

type OrderSuccessResponse struct {
	OrderID        string `json:"order_id"`
	ShipmentStatus string `json:"order_status"`
}
