package models

type User struct {
	ID       int		`json:id`
	login    string		`json:login`
	password string		`json:password`
	email 	 string		`json:email`
}