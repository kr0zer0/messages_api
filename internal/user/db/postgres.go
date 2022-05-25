package db

import (
	"example/messages_api/internal/user"
	"gorm.io/gorm"
)

type db struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) user.Repository {
	return &db{database: database}
}

func (d *db) Create(user user.User) {
	d.database.Create(&user)
}

func (d *db) GetByID(id string) (user user.User, err error) {
	err = d.database.Where("id = ?", id).First(&user).Error
	return user, err
}

func (d *db) GetByName(name string) (user user.User, err error) {
	err = d.database.Where("name = ?", name).First(&user).Error
	return user, err
}

func (d *db) GetAll() (users []user.User) {
	d.database.Find(&users)
	return users
}

func (d *db) Update(id string, updatedUser user.User) error {
	user, err := d.GetByID(id)
	if err != nil {
		return err
	}
	d.database.Model(&user).Updates(updatedUser)
	return nil
}

func (d *db) Delete(id string) error {
	user, err := d.GetByID(id)
	if err != nil {
		return err
	}
	d.database.Delete(&user)
	return nil
}
