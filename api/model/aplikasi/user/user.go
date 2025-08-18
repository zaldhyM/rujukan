package user

import (
	"log"
	"rujukan/model"
)

func QueryAll() ([]User, error) {
	var user []User
	if err := model.DB.Find(&user).Error; err != nil {
		log.Fatal(err)
	}

	return user, nil
}
