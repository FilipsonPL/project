package controllers

import (
	"github.com/FilipsonPL/project/models"
	"github.com/FilipsonPL/project/db"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user_data models.CreateUserInput

	if err := c.BindJSON(&user_data); err != nil{
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	user := models.User{Name:user_data.Name, Surname: user_data.Surname, Age: user_data.Age, Email: user_data.Email, Password:user_data.Password}
	db.Database.Create(&user)
	c.JSON(200, user)
}

func FindUsers(c *gin.Context) {
	var users []models.User
	db.Database.Find(&users)
	c.JSON(200,gin.H{"Users": users})
}

func FindUser(c *gin.Context) {
	var user models.User
	if err := db.Database.Where("id = ?").First(&user).Error; err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,gin.H{"User": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := db.Database.Where("id = ?",c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Database.Delete(&user)
	c.JSON(200,gin.H{"User deleted succesfully!":user})
}

func UpdateUser(c *gin.Context) {
	var user_data models.UpdateUser
	if err := c.BindJSON(&user_data); err != nil {
		c.JSON(400,gin.H{"Error":err.Error()})
		return
	}
	var user models.User
	if err := db.Database.Where("id = ?",c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	db.Database.Model(&user).Updates(&user_data)
	c.JSON(200,gin.H{"Edited user":user})
}
