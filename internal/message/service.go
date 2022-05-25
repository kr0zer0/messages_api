package message

import (
	"example/messages_api/internal/user"
	"example/messages_api/internal/user/db"
)

type service struct {
	repository Repository
}

func NewService(repo Repository) service {
	return service{repository: repo}
}

func (s *service) CreateMessage(input CreateMessageInput) error {
	userService := user.NewService(db.NewRepository(s.repository.GetDb()))

	sender, err := userService.GetUserByName(input.SenderName)
	if err != nil {
		return err
	}
	receiver, err := userService.GetUserByName(input.ReceiverName)
	if err != nil {
		return err
	}

	newMessage := Message{SenderID: sender.ID, ReceiverID: receiver.ID, MessageBody: input.MessageBody}
	s.repository.Create(newMessage)
	return nil
}

func (s *service) GetMessages() []Message {
	return s.repository.GetAll()
}

func (s *service) EditMessage(id string, editInput EditMessageInput) error {
	message, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}
	if editInput.MessageBody != "" {
		message.MessageBody = editInput.MessageBody
	}
	err = s.repository.Update(id, message)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteMessage(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
