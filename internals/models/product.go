package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model // This part will add automatically ID's, created_at, updated_at columons in the table created on mysql
	Name string
	Price int
}