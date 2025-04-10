package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Email	  string `json:"email" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Villages []Village `json:"villages" gorm:"foreignKey:UserID"`
}