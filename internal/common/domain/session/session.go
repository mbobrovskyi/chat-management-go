package session

import "github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/entity"

type Session interface {
	entity.Entity[Session]
}
