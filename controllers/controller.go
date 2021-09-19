package controllers

import (
	"github.com/FilipsonPL/project/models"
	"github.com/FilipsonPL/project/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var data map[string]string
	if err := c.BindJSON(&data); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	valid := true
	cells := []string{"name","surname","email","password"}
	for _, cell := range cells {
		if _, ok := data[cell]; !ok {
			valid = false
			break
		}
	}
	if valid {
		user := models.User{
			Name: data["name"],
			Surname: data["surname"],
			Email: data["email"],
			Password: data["password"],
		}
		db.Database.Create(&user)

		c.JSON(http.StatusOK, gin.H{"Created user": user})
	} else {
		c.JSON(http.StatusBadRequest,gin.H{"Error": "Not all values passed!"})
	}
}

func ShowUsers(c *gin.Context) {
	var users []models.User
	db.Database.Find(&users)
	c.JSON(http.StatusOK,gin.H{"Users": users})
}

func Unregister(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	var user models.User
	if _,ok := data["id"] ; ok {
		if err := db.Database.Where("id = ?",data["id"]).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Database.Delete(&user)
		c.JSON(http.StatusOK,gin.H{"User deleted succesfully!",user})
	} else {
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Missing id parameter!"})
	}
}

func EditUser(c *gin.Context) {
	var data map[string]string
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	var user models.User
	if _, ok := data["id"] ; ok {
		if err := db.Database.Where("id = ?",data["id"]).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for index, value := range data {
			if index!="id" {
				user[index] = value
			}
		}
		db.Database.Save(&user)
		c.JSON(http.StatusOK,gin.H{"Edited user":user})
		return
	} else {
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Missing id parameter!"})
		return
	}
}