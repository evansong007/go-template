package models

type Customer struct {
	CustomerID string `json:"customerId" gorm:"primaryKey"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}
