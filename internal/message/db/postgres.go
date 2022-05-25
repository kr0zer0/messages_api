package db

import (
	"example/messages_api/internal/message"
	"gorm.io/gorm"
)

type db struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) message.Repository {
	return &db{database: database}
}

func (d *db) GetDb() *gorm.DB {
	return d.database
}

func (d *db) Create(message message.Message) {
	d.database.Create(&message)
}

func (d *db) GetByID(id string) (message message.Message, err error) {
	err = d.database.Where("id = ?", id).First(&message).Error
	return message, err
}

func (d *db) GetAll() (messages []message.Message) {
	d.database.Find(&messages)
	return messages
}

func (d *db) Update(id string, updatedMessage message.Message) error {
	message, err := d.GetByID(id)
	if err != nil {
		return err
	}
	d.database.Model(message).Updates(updatedMessage)
	return nil
}

func (d *db) Delete(id string) error {
	message, err := d.GetByID(id)
	if err != nil {
		return err
	}
	d.database.Delete(&message)
	return nil
}
