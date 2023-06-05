package handler

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"time"
)

var JwtKey = jwt.NewHS256([]byte("your key"))

func LogIn(c *gin.Context) {
	// Verify username and password.
	// ...

	// If this user logged in successfully, give him/her a new JWT
	// Remember to assign value to subject header.
	payload := jwt.Payload{
		Subject:        "alice",
		ExpirationTime: jwt.NumericDate(time.Now().Add(time.Hour)),
	}
	token, err := jwt.Sign(payload, JwtKey)
	if err != nil {
		c.JSON(500, "some error")
		return
	}
	c.JSON(200, string(token))
}

func ReadBook(c *gin.Context) {
	c.String(200, "you read the book successfully")
}

func ReadAndWriteAndSoOn(c *gin.Context) {
	c.String(200, "you posted a book successfully")
}
