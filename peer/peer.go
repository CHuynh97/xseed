package peer

import (
	"github.com/google/uuid"
)

type Peer struct {
	ID        uuid.UUID
	SessionID uuid.UUID
	Peers     []*Peer
}

type PeerData map[string]interface{}

func (p Peer) Broadcast(PeerData) error
