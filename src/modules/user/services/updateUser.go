package services

import (
	"log"

	"boilerplate/src/modules/user/domain/entity"
	"boilerplate/src/modules/user/repository"
)

func UpdateUser(id string, updateUser *entity.User) *entity.User {

	user, err := repository.FindById(id)

	if err != nil {
		log.Panic("User not found")
	}

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.Password = updateUser.Password
	user.Active = updateUser.Active
	user.IsAdmin = updateUser.IsAdmin

	result, err := repository.Update(user)

	if err != nil {
		log.Panic("Erro on update account")
	}

	return result
}
