package services

import (
	"realmrovers/model"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (u *UserService) GetUsers(id uint) (*model.User, error) {
	var user model.User
	err := u.Db.Take(&user, id).Error
	if(err != nil ){
		return nil,err
	}
	return &user,nil
}
