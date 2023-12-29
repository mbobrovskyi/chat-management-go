package repositories

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/chat"
	"github.com/samber/lo"
	"time"
)

type ChatRepository struct {
	chats []chat.Chat
}

func (r *ChatRepository) getLastId() uint64 {
	return lo.MaxBy(r.chats, func(a chat.Chat, b chat.Chat) bool {
		return a.GetId() > b.GetId()
	}).GetId()
}

func (r *ChatRepository) GetAll(ctx context.Context) ([]chat.Chat, uint64, error) {
	return r.chats, uint64(len(r.chats)), nil
}

func (r *ChatRepository) GetById(ctx context.Context, id uint64) (chat.Chat, error) {
	for _, c := range r.chats {
		if c.GetId() == id {
			return chat.Create(c.GetId(), c.GetName(), c.GetType(), c.GetImage(), c.GetLastMessage(),
				c.GetCreatedBy(), c.GetCreatedAt(), c.GetUpdatedAt()), nil
		}
	}

	return nil, nil
}

func (r *ChatRepository) Save(ctx context.Context, c chat.Chat) (chat.Chat, error) {
	newChat := chat.Create(r.getLastId()+1, c.GetName(), c.GetType(), c.GetImage(), c.GetLastMessage(), c.GetCreatedBy(), c.GetCreatedAt(), c.GetUpdatedAt())

	r.chats = lo.Filter(r.chats, func(item chat.Chat, _ int) bool {
		return item.GetId() != newChat.GetId()
	})

	r.chats = append(r.chats, newChat)
	return c, nil
}

func (r *ChatRepository) Delete(ctx context.Context, id uint64) error {
	r.chats = lo.Filter(r.chats, func(item chat.Chat, _ int) bool {
		return item.GetId() != id
	})
	return nil
}

func NewChatRepository() chat.Repository {
	chats := make([]chat.Chat, 0)
	chats = append(chats, chat.Create(1, "Chat 1", chat.Direct, "", nil, 1, time.Now(), time.Now()))

	return &ChatRepository{
		chats: chats,
	}
}
