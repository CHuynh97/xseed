package peer

import "github.com/google/uuid"

type SessionPeer struct {
	SessionID uuid.UUID `db:"session_id"`
	PeerID    uuid.UUID `db:"peer_id"`
}
