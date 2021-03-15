package domain

import (
	"context"
	"time"
)

type Room struct {
	ID              int
	Name            string
	BackgroundURL   string
	BackgroundColor string
	UserCount       int
	CreatedAt       time.Time
}

var _ INode = (*Room)(nil)

func (r *Room) GetID() int {
	return r.ID
}

func (r *Room) GetCreatedAtUnix() int {
	return int(r.CreatedAt.Unix())
}

type RoomListReq struct {
	Limit         int
	LastKnownID   int
	LastKnownUnix int
}

type RoomListResp struct {
	List    []*Room
	HasNext bool
}

type IRoomRepo interface {
	GetByID(ctx context.Context, id int) (*Room, error)
	List(ctx context.Context, req *RoomListReq) (*RoomListResp, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, room *Room) error
}
