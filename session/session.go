package session

import (
	"context"

	proto "github.com/CHuynh97/xseed/api/gen"
	internal "github.com/CHuynh97/xseed/internal/session"
	"github.com/google/uuid"
)

type SessionID uuid.UUID

type SessionService struct {
	internal.SessionStorage
}

var _ proto.SessionServiceServer = SessionService{}

func (s SessionService) Create(ctx context.Context, pb *proto.CreateSessionRequest) (*proto.CreateSessionResponse, error)

func (s SessionService) Destroy(ctx context.Context, pb *proto.DestroySessionRequest) (*proto.DestroySessionResponse, error)

func (s SessionService) Join(ctx context.Context, pb *proto.JoinSessionRequest) (*proto.JoinSessionResponse, error)

func NewSessionService() (*SessionService, error) {
	return &SessionService{}, nil
}
