package session

import (
	"time"

	"github.com/CHuynh97/xseed/state"
	"github.com/google/uuid"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Session struct {
	ID        uuid.UUID `db:"id"`
	StateID   uuid.UUID `db:"state_id"`
	Active    bool      `db:"active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type SessionStorage struct {
	db postgresql.ConnectionURL
}

func (s SessionStorage) CreateSession(state state.StateData) (uuid.UUID, error)

func (s SessionStorage) DeleteSession(uuid.UUID) error
