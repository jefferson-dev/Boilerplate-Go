package services

import (
	"log"

	"boilerplate/src/modules/user/domain/entity"
	"boilerplate/src/modules/user/repository"
)

func DeleteUser(id string) error {
	err := repository.Delete(id, &entity.User{})

	if err != nil {
		log.Panic("User not found")
	}

	return nil
}
