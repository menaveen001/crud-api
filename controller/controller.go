package controller

import (
	"curd-api/database"
	"curd-api/entity"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// insertdata into db
func InsertUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ab entity.User
		if err := c.BindJSON(&ab); err != nil {
			log.Printf("Invalid Request JSON %v", err)
		}

		result := database.InsertUser(ab)
		c.JSON(200, result)
	}

}

// get all data from db
func GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, database.SelectData())
	}

}

// get data by id from db

func GetDataByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ab entity.User
		id := c.Param("id")
		nid, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err, nid, ab)
		}

		c.JSON(200, database.SelectDataByID(nid))
	}

}

// delete data by id from db

func DeleteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ab entity.User
		id := c.Param("id")
		c.BindJSON(&ab)
		nid, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err, nid)

		}
		c.JSON(200, database.DeleteDataByID(nid))
		log.Println("data is deleted")
	}

}

// update data by id into db
func UpdateById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ab entity.User

		if err := c.BindJSON(&ab); err != nil {
			log.Printf("Invalid Request JSON %v", err)
		}
		c.JSON(200, database.UpdateUserById(ab))
		log.Println(ab)
	}

}
