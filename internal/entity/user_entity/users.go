package user_entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PhoneNumber string
	FacebookId  string
	LoginType   string
	Status      bool
}

func (User) TableName() string {
	return "public.users"
}
