package entity_test

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

type testEntity struct {
	entity.Entity[testEntity]
}

func newTestEntity(id uint64) *testEntity {
	return &testEntity{
		Entity: entity.New[testEntity](id),
	}
}

func TestNewEntityId(t *testing.T) {
	_entity := newTestEntity(1)
	require.EqualValues(t, 1, _entity.GetId())
}

func TestNewEntity_Equal(t *testing.T) {
	_entity1 := newTestEntity(1)
	_entity2 := newTestEntity(1)
	require.True(t, _entity1.Equals(_entity2))
}

func TestNewEntity_NotEqual(t *testing.T) {
	_entity1 := newTestEntity(1)
	_entity2 := newTestEntity(2)
	require.False(t, _entity1.Equals(_entity2))
}
