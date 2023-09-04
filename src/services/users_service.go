package services

import "rock/test_gin/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
