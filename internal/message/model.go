package message

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID    uint   `json:"sender_id"`
	ReceiverID  uint   `json:"receiver_id"`
	MessageBody string `json:"message_body"`
}

type CreateMessageInput struct {
	SenderName   string `json:"sender_name"`
	ReceiverName string `json:"receiver_name"`
	MessageBody  string `json:"message_body"`
}

type EditMessageInput struct {
	MessageBody string `json:"message_body"`
}
