package handler

import (
	"casbin-go_gin/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := h.Services.Auth.CreateUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, input)
}

func (h *Handler) LogIn(c *gin.Context) {
	var inputUser models.User
	if err := c.BindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	token, err := h.Services.Auth.GenerateToken(inputUser.Name, inputUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (h *Handler) ReadBook(c *gin.Context) {
	c.String(200, "you read the book successfully")
}

func (h *Handler) ReadAndWriteAndSoOn(c *gin.Context) {
	c.String(200, "you posted a book successfully")
}
