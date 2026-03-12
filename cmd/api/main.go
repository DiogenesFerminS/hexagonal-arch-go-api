package main

import (
	"fmt"
	"go-api/cmd/api/db"
	"go-api/cmd/api/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Load environment variables failed")
		return
	}

	db.InitDB()

	server := gin.Default()

	server.POST("/users", func(ctx *gin.Context) {
		var user User

		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ok":      false,
				"message": "Bad Request",
				"error":   "Invalid data",
			})
			return
		}

		query := "INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id"

		hashedPassword, err := utils.HashPassword(user.Password)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ok":      false,
				"message": "Bad Request",
				"error":   "Hashed password failed",
			})
			return
		}

		err = db.DB.QueryRow(query, user.Username, hashedPassword).Scan(&user.ID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"ok":      false,
				"message": "Bad Request",
				"error":   "Query failed",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Success",
			"data": gin.H{
				"id":       user.ID,
				"username": user.Username,
			},
		})
	})

	port := os.Getenv("PORT")
	err = server.Run(":" + port)

	if err != nil {
		fmt.Println("Run server failed" + err.Error())
	}
}
