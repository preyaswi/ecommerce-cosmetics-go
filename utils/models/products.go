package models

type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category_name"`
}
type ProductBrief struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	SKU           string  `json:"sku"`
	CategoryName  string  `json:"category"`
	BrandID       uint    `json:"brand_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	ProductStatus string  `json:"product_status"`
}
type ProductReceiver struct {
	Name              string  `json:"name"`
	SKU               string  `json:"sku"`
	CategoryId        string  `json:"category_id"`
	DesignDescription string  `json:"design_description"`
	BrandID           uint    `json:"brand_id"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
	ProductStatus     string  `json:"product_status"`
}
type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type ProductUpdate struct {
	ProductId int `json:"product_id"`
	Quantity     int `json:"quantity"`
}
type ProductUpdateReciever struct {
	ProductID int
	Quantity     int
}
