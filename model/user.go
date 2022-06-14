package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (u *User) Create() {
	// result := db.Create(&user)
}
