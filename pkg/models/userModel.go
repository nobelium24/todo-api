package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `grom:"unique" json:"email"`
	Password  string `json:"password"`
	Todos     []Todo `json:"todos"`
}
