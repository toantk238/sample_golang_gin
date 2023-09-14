package services

import (
	"rock/test_gin/domain/users"
	"rock/test_gin/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
