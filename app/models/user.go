package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email string `gorm:"uniqueIndex"`
	Password string
}
