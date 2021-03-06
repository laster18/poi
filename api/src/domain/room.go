package domain

import (
	"context"
	"time"

	val "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/laster18/poi/api/src/util/clock"
	"github.com/laster18/poi/api/src/util/validator"
)

type Room struct {
	ID              int
	Name            string
	BackgroundURL   string
	BackgroundColor string
	UserCount       int
	CreatedAt       time.Time
}

const (
	DefaultBgColor = "#20b2aa"
	DefaultBgURL   = "https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg1.jpg"
)

func NewRoom(name string, bgColor, bgURL *string) *Room {
	color := DefaultBgColor
	if bgColor != nil {
		color = *bgColor
	}

	url := DefaultBgURL
	if bgURL != nil {
		url = *bgURL
	}

	return &Room{
		ID:              0,
		Name:            name,
		BackgroundURL:   url,
		BackgroundColor: color,
		UserCount:       0,
		CreatedAt:       clock.Now(),
	}
}

var _ INode = (*Room)(nil)

func (r *Room) GetID() int {
	return r.ID
}

func (r *Room) GetCreatedAtUnix() int {
	return int(r.CreatedAt.Unix())
}

func (r *Room) Validate() error {
	return validator.ValidateStruct(r,
		val.Field(&r.Name, val.Required, val.RuneLength(2, 20)),
		val.Field(&r.BackgroundColor, val.Required, is.HexColor),
		val.Field(&r.BackgroundURL, val.Required, is.URL),
	)
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
	GetByName(ctx context.Context, name string) (*Room, error)
	List(ctx context.Context, req *RoomListReq) (*RoomListResp, error)
	ListAll(ctx context.Context) ([]*Room, error)
	Create(ctx context.Context, room *Room) error
	Count(ctx context.Context) (int, error)
}
