package domain

import "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/chat"

type ChatService interface {
	GetAll() ([]chat.Chat, error)
	GetById(id uint64) ([]chat.Chat, error)
	Save(chat chat.Chat) (chat.Chat, error)
	Delete(id uint64) error
}

type chatService struct {
	repository Repository
}

func (s *chatService) GetAll() ([]chat.Chat, error) {
	return s.repository.GetAll()
}

func (s *chatService) GetById(id uint64) ([]chat.Chat, error) {
	return s.repository.GetAll()
}

func (s *chatService) Save(chat chat.Chat) (chat.Chat, error) {
	return s.repository.Save(chat)
}

func (s *chatService) Delete(id uint64) error {
	return s.repository.Delete(id)
}

func NewService(repository Repository) ChatService {
	return &chatService{repository: repository}
}
