package discovery

import (
	"github.com/CHuynh97/xseed/peer"
	"github.com/CHuynh97/xseed/session"
	"github.com/upper/db/v4/adapter/postgresql"
)

type DiscoveryStorage struct {
	db postgresql.ConnectionURL
}

func (d DiscoveryStorage) ListPeers(sessionID session.SessionID) ([]peer.Peer, error)
