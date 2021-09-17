package main

import (
	"database/sql"
	//"github.com/gin-gonic/gin"
	//"net/http"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID       int
	login    string
	password string
	email string
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
		panic("Exit")
	}
}

func LoadDataFromDatabase(sqlObject *sql.DB) []User {
	var table []User
	rows,err := sqlObject.Query("SELECT * from user")
	checkError(err)
	var temp User
	for rows.Next() {
		rows.Scan(&temp.ID,&temp.login,&temp.password,&temp.email)
		table = append(table,temp)
	}
	return table
}

func AddNewUser(sqlObject *sql.DB) {
	input,err := sqlObject.Prepare("INSERT INTO user (login, password, email) VALUES(?,?,?)")
	checkError(err)
	var temp User
	fmt.Scanf("%s %s %s",&temp.login,&temp.password,&temp.email)
	result,err := input.Exec(temp.login,temp.password,&temp.email)
	checkError(err)
	fmt.Printf("%v\n",result)
}

func EditUser(sqlObject *sql.DB,id int) {

}

func main() {
	dataBase,err := sql.Open("sqlite3","users.db")
	checkError(err)
	
	users := LoadDataFromDatabase(dataBase)
	fmt.Printf("%v\n",users)
	//AddNewUser(dataBase)

	//ser := gin.Default()
	//ser.POST("/add",postAddUser)
	//ser.POST("/delete",postDeleteUser)
	//ser.POST("/Modify",postModifyUser)

	//ser.Run("localhost:8080")*/
}
