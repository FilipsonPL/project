package models

type User struct {
	ID        	uint64 `json:"id" gorm:"primary_key"`
	Name 		string `json:"name"`
	Surname  	string `json:"surname"`
	Email		string `json:"email" gorm:"unique"`
	Password    string `json:"password" `
}