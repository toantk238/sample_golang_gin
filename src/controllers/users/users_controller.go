package users

import (
	"fmt"
	"net/http"
	"rock/test_gin/domain/users"
	"rock/test_gin/services"

	"github.com/gin-gonic/gin"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
    //TODO: return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
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
