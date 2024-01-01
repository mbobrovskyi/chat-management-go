package repositories

import (
	"context"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/abstracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/valueobjects"
	"github.com/samber/lo"
	"time"
)

var _ abstracts.ChatRepository = (*MemoryChatRepository)(nil)

type MemoryChatRepository struct {
	chats []chat.Chat
}

func (r *MemoryChatRepository) getLastId() uint64 {
	if len(r.chats) == 0 {
		return 0
	}

	found := lo.MaxBy(r.chats, func(left chat.Chat, right chat.Chat) bool {
		return left.Id() > right.Id()
	})

	return found.Id()
}

func (r *MemoryChatRepository) GetAll(ctx context.Context) ([]chat.Chat, uint64, error) {
	return r.chats, uint64(len(r.chats)), nil
}

func (r *MemoryChatRepository) GetById(ctx context.Context, id uint64) (*chat.Chat, error) {
	for _, chat := range r.chats {
		if chat.Id() == id {
			return &chat, nil
		}
	}

	return nil, nil
}

func (r *MemoryChatRepository) Save(ctx context.Context, c chat.Chat) (*chat.Chat, error) {
	var newChat chat.Chat

	if c.Id() == 0 {
		newChat = chat.New(
			r.getLastId()+1, c.Name(), c.Description(), c.Type(), c.Image(), c.LastMessage(),
			c.MemberIds(), c.CreatedBy(), c.CreatedAt(), c.UpdatedAt(),
		)
	} else {
		newChat = c
	}

	r.chats = lo.Filter(r.chats, func(item chat.Chat, _ int) bool {
		return item.Id() != newChat.Id()
	})

	r.chats = append(r.chats, newChat)
	return &newChat, nil
}

func (r *MemoryChatRepository) Delete(ctx context.Context, id uint64) error {
	r.chats = lo.Filter(r.chats, func(item chat.Chat, _ int) bool {
		return item.Id() != id
	})
	return nil
}

func NewMemoryChatRepository() *MemoryChatRepository {
	chat1 := chat.New(1, "Chat 1", "", valueobjects.Direct, "", nil, []uint64{1, 2}, 1, time.Now(), time.Now())
	return &MemoryChatRepository{chats: []chat.Chat{chat1}}
}
