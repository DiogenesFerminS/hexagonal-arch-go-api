package user

import (
	"go-api/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateUser(ctx *gin.Context) {
	var userDto domain.User

	err := ctx.ShouldBindJSON(&userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Bad Request",
			"error":   "Invalid data",
		})
		return
	}

	userId, err := h.UserService.Create(userDto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Success",
		"data": gin.H{
			"id": userId,
		},
	})
}
