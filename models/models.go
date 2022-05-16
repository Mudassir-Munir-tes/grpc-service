package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
}

type Driver struct {
	gorm.Model
	Name string
}

type Truck struct {
	gorm.Model
	ModelNo int32
	Power   string
}
