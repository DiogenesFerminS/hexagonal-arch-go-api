package main

import (
	"fmt"
	"go-api/cmd/api/db"
	"go-api/internal/services/user"

	"os"

	postgresRepo "go-api/internal/repositories/postgresql/user"

	userHandler "go-api/cmd/api/handlers/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Load environment variables failed")
		return
	}

	db.InitDB()
	defer db.DB.Close()

	server := gin.Default()

	postgresRepo := postgresRepo.NewResposity(db.DB)

	userService := user.UserService{
		Repository: postgresRepo,
	}

	usersHandler := userHandler.Handler{
		UserService: userService,
	}

	server.POST("/users", usersHandler.CreateUser)

	port := os.Getenv("PORT")
	err = server.Run(":" + port)

	if err != nil {
		fmt.Println("Run server failed" + err.Error())
	}
}
