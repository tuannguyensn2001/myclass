package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int            `db:"id" gorm:"column:id;" json:"id"`
	Username  string         `db:"username" gorm:"column:username;" json:"username"`
	Password  string         `db:"password" gorm:"column:password;" json:"password"`
	Email     string         `db:"email" gorm:"column:email;" json:"email"`
	CreatedAt time.Time      `db:"created_at" gorm:"column:created_at;"  json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"column:updated_at"  json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
