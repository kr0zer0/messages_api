package message

type CreateMessageInput struct {
	SenderName   string `json:"sender_name"`
	ReceiverName string `json:"receiver_name"`
	MessageBody  string `json:"message_body"`
}

type EditMessageInput struct {
	MessageBody string `json:"message_body"`
}
