package discovery

import (
	"context"

	proto "github.com/CHuynh97/xseed/api/gen"
)

type DiscoveryService struct{}

var _ proto.DiscoveryServiceServer = DiscoveryService{}

func (d DiscoveryService) ListPeers(ctx context.Context, pb *proto.ListPeersRequest) (*proto.ListPeersResponse, error)
