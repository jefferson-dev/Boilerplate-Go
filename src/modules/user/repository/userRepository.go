package repository

import (
	"boilerplate/src/infra/database"
	"boilerplate/src/modules/user/domain/entity"
)

func GetAll() (*[]entity.User, error) {
	user := []entity.User{}

	err := database.DB.Find(&user)

	return &user, err.Error
}

func FindById(id string) (*entity.User, error) {
	user := entity.User{}

	err := database.DB.First(&user, "id = ?", id)

	return &user, err.Error
}

func Create(user *entity.User) (*entity.User, error) {
	err := database.DB.Create(user)

	return user, err.Error
}

func Update(user *entity.User) (*entity.User, error) {

	err := database.DB.Save(&user)

	return user, err.Error
}

func Delete(id string, user *entity.User) error {

	err := database.DB.Delete(&user, "id = ?", id)

	return err.Error
}
