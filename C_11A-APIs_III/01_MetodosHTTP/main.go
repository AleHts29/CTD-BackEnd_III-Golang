package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserPatchReq struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

var Users = []User{
	{1, "Juan Pablo", "juan@fakemail.com"},
	{2, "José Perez", "jose@fakemail.com"},
	{3, "Jaime Pogo", "jaime@fakemail.com"},
	{4, "camilo Ruiz", "camilo@fakemail.com"},
}

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, Users)
		return
	})

	r.POST("/login", func(c *gin.Context) {
		var login Login
		c.BindJSON(&login) //capturamos lo que viene del body(postman) y populamos la variable login
		c.JSON(200, gin.H{"status": login.Email})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		c.BindJSON(&user) //capturamos lo que viene del body(postman) y populamos la variable user
		idReceived := c.Param("id")
		idReceivedInt, err := strconv.Atoi(idReceived)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, u := range Users {
			if u.ID == idReceivedInt {
				user.ID = idReceivedInt
				Users[i] = user
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", idReceivedInt)})
		return
	})

	r.PATCH("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var userReq UserPatchReq
		err = c.BindJSON(&userReq) //capturamos lo que viene del body(postman) y populamos la variable userReq
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, user := range Users {
			if user.ID == userIDInt {
				if userReq.ID != nil {
					user.ID = *userReq.ID
				}
				if userReq.Username != nil {
					user.Username = *userReq.Username
				}
				if userReq.Email != nil {
					user.Email = *userReq.Email
				}
				Users[i] = user
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", userIDInt)})
		return
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, user := range Users {
			if user.ID == userIDInt {
				Users = append(Users[:i], Users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("user %d deleted", userIDInt)})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", userIDInt)})
		return
	})

	r.Run(":8080")
}
