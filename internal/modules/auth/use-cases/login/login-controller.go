package login

import (
	login "go-api/internal/modules/auth/use-cases/login/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var dto login.LoginDto

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginService(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
	c.Set("Authorization", token)

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
