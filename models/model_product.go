package models

type Product struct {
	ProductID string  `json:"productId" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Price     float64 `json:"price" gorm:"type:numeric"`
	VendorID  string  `json:"vendorId"`
}
