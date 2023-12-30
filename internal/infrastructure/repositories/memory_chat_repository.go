package repositories

import (
	"context"
	chat2 "github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/interfaces"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/valueobjects"
	"github.com/samber/lo"
	"time"
)

var _ interfaces.ChatRepository = (*MemoryChatRepository)(nil)

type MemoryChatRepository struct {
	chats []chat2.Chat
}

func (r *MemoryChatRepository) getLastId() uint64 {
	if len(r.chats) == 0 {
		return 0
	}

	found := lo.MaxBy(r.chats, func(left chat2.Chat, right chat2.Chat) bool {
		return left.Id() > right.Id()
	})

	return found.Id()
}

func (r *MemoryChatRepository) GetAll(ctx context.Context) ([]chat2.Chat, uint64, error) {
	return r.chats, uint64(len(r.chats)), nil
}

func (r *MemoryChatRepository) GetById(ctx context.Context, id uint64) (*chat2.Chat, error) {
	for _, chat := range r.chats {
		if chat.Id() == id {
			return &chat, nil
		}
	}

	return nil, nil
}

func (r *MemoryChatRepository) Save(ctx context.Context, c chat2.Chat) (*chat2.Chat, error) {
	var newChat chat2.Chat

	if c.Id() == 0 {
		newChat = chat2.New(
			r.getLastId()+1, c.Name(), c.Description(), c.Type(), c.Image(), c.LastMessage(),
			c.MemberIds(), c.CreatedBy(), c.CreatedAt(), c.UpdatedAt(),
		)
	} else {
		newChat = c
	}

	r.chats = lo.Filter(r.chats, func(item chat2.Chat, _ int) bool {
		return item.Id() != newChat.Id()
	})

	r.chats = append(r.chats, newChat)
	return &newChat, nil
}

func (r *MemoryChatRepository) Delete(ctx context.Context, id uint64) error {
	r.chats = lo.Filter(r.chats, func(item chat2.Chat, _ int) bool {
		return item.Id() != id
	})
	return nil
}

func NewMemoryChatRepository() *MemoryChatRepository {
	chat1 := chat2.New(1, "Chat 1", "", valueobjects.Direct, "", nil, []uint64{1, 2}, 1, time.Now(), time.Now())
	return &MemoryChatRepository{chats: []chat2.Chat{chat1}}
}
