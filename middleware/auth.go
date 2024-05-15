package middleware

import (
	"database/sql"
	"errors"
	"halosus/db"
	"halosus/helper/jwt"
	"halosus/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func getBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func AuthMiddleware(c *gin.Context) {
	token, err := getBearerToken(c.GetHeader("Authorization"))
	if err!= nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": err.Error()})
		return
	}
	id, err := jwt.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message":err.Error()})
		return
	}
	// find user
	db := db.GetDB()
	var user models.User
	err = db.QueryRow("SELECT id, nip FROM public.user WHERE id = ? LIMIT 1",id).Scan(&user.Id, &user.Nip)
	if err != nil {
		if err == sql.ErrNoRows{
			c.AbortWithStatusJSON(404, gin.H{
				"message":"user not found"})
				return
		}
	}
	c.Set("userId",id)
	c.Next()
}