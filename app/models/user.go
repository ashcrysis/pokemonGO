package models

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Email       string `gorm:"uniqueIndex"`
	Name        string 
	Phone       string 
	PostalCode  string 
	Street      string 
	Number      string
	Complement  string 
	Password    string
	Image_url string
}