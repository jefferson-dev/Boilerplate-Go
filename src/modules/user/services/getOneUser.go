package services

import (
	"log"

	"boilerplate/src/modules/user/domain/entity"
	"boilerplate/src/modules/user/repository"
)

func GetOneUser(id string) *entity.User {
	result, err := repository.FindById(id)

	if err != nil {
		log.Panic("User not found")
	}

	return result
}
