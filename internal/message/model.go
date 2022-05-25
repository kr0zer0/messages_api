package message

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID    uint   `json:"sender_id"`
	ReceiverID  uint   `json:"receiver_id"`
	MessageBody string `json:"message_body"`
}
