package services

import (
	"log"

	"boilerplate/src/modules/user/domain/entity"
	"boilerplate/src/modules/user/repository"
)

func CreateUser(user *entity.User) *entity.User {
	result, err := repository.Create(user)

	if err != nil {
		log.Panic("Erro on create account")
	}

	return result
}
