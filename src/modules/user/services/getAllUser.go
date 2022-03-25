package services

import (
	"log"

	"boilerplate/src/modules/user/domain/entity"
	"boilerplate/src/modules/user/repository"
)

func GetAllUser() *[]entity.User {
	result, err := repository.GetAll()

	if err != nil {
		log.Panic("User not found")
	}

	return result
}
