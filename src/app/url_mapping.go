package app

import (
	"rock/test_gin/controllers/ping"
	"rock/test_gin/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
