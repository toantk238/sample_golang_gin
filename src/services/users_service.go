package services

import (
	"net/http"
	"rock/test_gin/domain/users"
	"rock/test_gin/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
  return nil, nil
	
  return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
