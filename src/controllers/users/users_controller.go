package users

import (
	"net/http"
	"rock/test_gin/domain/users"
	"rock/test_gin/services"
	"rock/test_gin/utils/errors"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me !")

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me !")
}
