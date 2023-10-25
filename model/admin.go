package model

type Admin struct {
	ID			uint 	`gorm:"PrimaryKey"`
	Name 		string
	Email		string 	`gorm:"Unique"`
	Password	string
}