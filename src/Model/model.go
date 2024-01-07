package model

import (
    "gorm.io/gorm"
)

type Dto struct {
	gorm.Model
	Name string 
	Id int `gorm:"primaryKey"`
	Password string
	Email string `gorm:"unique"`
}