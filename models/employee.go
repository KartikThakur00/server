package models

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	FromDate time.Time `json:"from"`
	ToDate time.Time `json:"to"`
	Phone int64 `json:"phone"`
	Resume string `json:"resume"`
	Email string `json:"email"`
}