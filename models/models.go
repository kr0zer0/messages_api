package models

import "gorm.io/gorm"

// User represents the user model
type User struct {
	gorm.Model
	Name string `json:"name"`
}

// Message represents the message model
type Message struct {
	gorm.Model
	SenderID    uint   `json:"sender_id"`
	ReceiverID  uint   `json:"receiver_id"`
	MessageBody string `json:"message_body"`
}

// CreateMessageInput operates request body when creating a new message
type CreateMessageInput struct {
	SenderName   string `json:"sender_name"`
	ReceiverName string `json:"receiver_name"`
	MessageBody  string `json:"message_body"`
}

// EditMessageInput operates request body when editing the message
type EditMessageInput struct {
	MessageBody string `json:"message_body"`
}
