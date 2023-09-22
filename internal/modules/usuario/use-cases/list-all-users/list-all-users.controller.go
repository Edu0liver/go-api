package list_all_users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllUsersController(c *gin.Context) {
	if user, err := ListAllUsersService(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}
