package message

import "gorm.io/gorm"

type Repository interface {
	GetDb() *gorm.DB
	Create(message Message)
	GetByID(id string) (Message, error)
	GetAll() []Message
	Update(id string, updatedMessage Message) error
	Delete(id string) error
}
