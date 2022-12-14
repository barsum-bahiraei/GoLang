package Models

import "time"

type User struct {
	ID          int    `gorm:"primaryKey"`
	Username    string `gorm:"unique"`
	Password    string
	Name        string
	Family      string
	PhoneNumber string `gorm:"unique"`
	Age         int
	CreatedAt   time.Time
}
