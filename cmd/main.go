package main

import (
	"casbin-go_gin/handler"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"github.com/maxwellhertz/gin-casbin"
	"log"
	"strings"
)

func main() {
	r := gin.Default()

	r.POST("/login", handler.LogIn)

	// Use Casbin authentication middleware.
	auth, err := gcasbin.NewCasbinMiddleware("path/model.conf", "path/policy.csv", subjectFromJWT)
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/book", auth.RequiresPermissions([]string{"book:read"}, gcasbin.WithLogic(gcasbin.AND)), handler.ReadBook)

	r.POST("/book", auth.RequiresRoles([]string{"guest"}, gcasbin.WithLogic(gcasbin.AND)), handler.ReadAndWriteAndSoOn)

	if err = r.Run(":8888"); err != nil {
		log.Fatalln(err)
	}
}

// subjectFromJWT parses a JWT and extract subject from sub claim.
func subjectFromJWT(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	prefix := "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		// Incorrect Authorization header format.
		return ""
	}
	token := authHeader[strings.Index(authHeader, prefix)+len(prefix):]
	if token == "" {
		// JWT not found.
		return ""
	}

	var payload jwt.Payload
	_, err := jwt.Verify([]byte(token), handler.JwtKey, &payload)
	if err != nil {
		return ""
	}
	return payload.Subject
}
