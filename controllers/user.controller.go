package controllers

import (
	"database/sql"
	"halosus/db"
	"halosus/dto"
	"halosus/helper"
	"halosus/helper/hash"
	"halosus/helper/jwt"
	"halosus/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (h UserController) CreateIT(c *gin.Context) {
	db := db.GetDB()

	var request dto.CreateITRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// log.Println(request)
	if !helper.ValidateNip(request.Nip, 615) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid nip"})
		return
	}
	
	if !helper.ValidateString(request.Name, 5, 50) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid name"})
		return
	}
	
	if !helper.ValidateString(request.Password, 5, 33) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}
	
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM public.user WHERE nip = $1", request.Nip).Scan(&count)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count > 0 {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "nip already exists."})
		return
	}

	hashedPass, err := hash.HashPassword(request.Password)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request.Password = hashedPass

	var user models.User
	if err := db.QueryRow("INSERT INTO public.user (nip, name, password, \"createdAt\", \"updatedAt\") VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, nip, name", request.Nip, request.Name, request.Password).Scan(&user.Id, &user.Nip, &user.Name); err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.SignJWT(user.Id, user.Nip)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := map[string]any{
		"userId": user.Id,
		"nip": user.Nip,
		"name": user.Name,
		"accessToken": token,
	}
	payload := gin.H{
		"message": "User registered successfully",
		"data": data,
	}
	c.JSON(http.StatusCreated, payload)
}

func (h UserController) ITLogin(c *gin.Context) {
	db := db.GetDB()

	var request dto.LoginITRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !helper.ValidateNip(request.Nip, 615) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid nip"})
		return
	}
	if !helper.ValidateString(request.Password, 5, 33) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	var user models.User
	q := `
		SELECT id, nip, name, password FROM public.user WHERE nip = $1
	`
	err := db.QueryRow(q, request.Nip).Scan(&user.Id, &user.Nip, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !hash.CheckPassword(request.Password, user.Password){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wrong password"})
		return
	}

	token, err := jwt.SignJWT(user.Id, user.Nip)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := map[string]any{
		"userId": user.Id,
		"nip": user.Nip,
		"name": user.Name,
		"accessToken": token,
	}

	payload := gin.H{
		"message": "User logged in successfully",
		"data": data,
	}
	c.JSON(http.StatusOK, payload)
}