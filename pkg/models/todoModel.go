package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Task   string `json:"task"`
	Status string `json:"status"`
	UserId uint   `json:"userId"`
}
