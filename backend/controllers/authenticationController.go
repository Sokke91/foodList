package controllers

import (
	"foodList/authentication"
	"foodList/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var userInput dto.AuthenticationDto

	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	if userInput.Username != "Admin" || userInput.Password != "geheim" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token := authentication.GenerateToken(userInput.Username, ctx)

	ctx.JSON(http.StatusOK, gin.H{"message": token})
}

func ProtectedRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Login correct"})
}
