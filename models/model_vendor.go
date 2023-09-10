package models

type Vendor struct {
	VendorID string `json:"vendorId" gorm:"primaryKey"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Phone    string `json:"phoneNumber"`
	Email    string `json:"emailAddress"`
	Address  string `json:"address"`
}
