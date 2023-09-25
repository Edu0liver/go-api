package list_all_users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllUsersController(c *gin.Context) {
	user, err := ListAllUsersService()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
