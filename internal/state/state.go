package state

import (
	"github.com/CHuynh97/xseed/session"
	"github.com/CHuynh97/xseed/state"
	"github.com/redis/go-redis/v9"
)

type StateCache struct {
	redisClient redis.Client
}

func (s StateCache) LoadState(sessionID session.SessionID, state state.StateData) error

func (s StateCache) Gettate(sessionID session.SessionID) error
