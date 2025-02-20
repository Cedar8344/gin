package controllers

import (
	"log"
	"strconv"
	"fmt"
	"gin-login/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetUser(c *gin.Context) {
	fmt.Println("hello")
	var user []models.User
	_, err := dbmap.Select(&user, "select * from user")
	
	if err == nil {
		c.JSON(200, user)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found also fuck you"})
	}

}

func GetUserDetail(c *gin.Context) {
	fmt.Println("hello2")
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)

	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)

		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func Login(c *gin.Context) {
	fmt.Println("hello3")
	var user models.User
	c.Bind(&user)
	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)

	if err == nil {
		user_id := user.Id

		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func PostUser(c *gin.Context) {
fmt.Println("hello4")
	var user models.User
	c.Bind(&user)

	log.Println(user)
	d := user.Id+1
	log.Println(d)
	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {

		if insert, err := dbmap.Exec(`INSERT INTO user (Id, Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?, ?)`, user.Id, user.Username, user.Password, user.Firstname, user.Lastname); insert != nil {
			//user_id , err := insert.LastInsertId()
			if err == nil {
				
				content := &models.User{
					//Id: 	user_id,
					Id:      user.Id,
					Username:  user.Username,
					Password:  user.Password,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

}

func UpdateUser(c *gin.Context) {
fmt.Println("hello5")
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)

	if err == nil {
		var json models.User
		c.Bind(&json)

		user_id, _ := strconv.ParseInt(id, 0, 64)

		user := models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}

		if user.Firstname != "" && user.Lastname != "" {
			_, err = dbmap.Update(&user)

			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
