package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/middleware"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input *model.CreateRoomInput) (*model.CreateRoomPayload, error) {
	dupRoom, err := r.roomRepo.GetByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if dupRoom != nil {
		return nil, fmt.Errorf("%q is already exists", input.Name)
	}

	newRoom := domain.NewRoom(input.Name, "#20b2aa")
	if err := newRoom.Validate(); err != nil {
		return nil, err
	}

	if err := r.roomRepo.Create(ctx, newRoom); err != nil {
		return nil, err
	}

	return toCreateRoomPayload(newRoom), nil
}

func (r *queryResolver) Rooms(ctx context.Context, first *int, after *string, orderBy *model.RoomOrderField) (*model.RoomConnection, error) {
	roomListReq := &domain.RoomListReq{}

	if first != nil {
		roomListReq.Limit = *first
	}
	if after != nil {
		// afterCursorのformatチェック
		id, unix, err := decodeCursor(roomPrefix, after)
		if err != nil {
			return nil, err
		}

		roomListReq.LastKnownID = id
		roomListReq.LastKnownUnix = unix
	}

	roomListResp, err := r.roomRepo.List(ctx, roomListReq)
	if err != nil {
		return nil, err
	}
	roomCount, err := r.roomRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// create pageInfo
	var pageInfo model.PageInfo
	pageInfo.HasNextPage = roomListResp.HasNext
	if after != nil {
		pageInfo.HasPreviousPage = true
	} else {
		pageInfo.HasPreviousPage = false
	}
	startCursor, endCursor := getRoomCursors(roomListResp.List)
	pageInfo.StartCursor = startCursor
	pageInfo.EndCursor = endCursor

	// serialize room model
	nodes := make([]*model.Room, len(roomListResp.List))
	for i, room := range roomListResp.List {
		nodes[i] = &model.Room{
			ID:   strconv.Itoa(int(room.ID)),
			Name: room.Name,
		}
	}

	// create connection
	edges := make([]*model.RoomEdge, len(roomListResp.List))
	for i, room := range roomListResp.List {
		edges[i] = &model.RoomEdge{
			Cursor: *encodeCursor(roomPrefix, room.ID, int(room.CreatedAt.Unix())),
			Node:   nodes[i],
		}
	}

	return &model.RoomConnection{
		PageInfo:  &pageInfo,
		Nodes:     nodes,
		Edges:     edges,
		RoomCount: roomCount,
	}, nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	roomID, err := decodeID(roomPrefix, id)
	if err != nil {
		return nil, err
	}

	room, err := r.roomRepo.GetByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return &model.Room{
		ID:   strconv.Itoa(room.ID),
		Name: room.Name,
	}, nil
}

func (r *roomResolver) ID(ctx context.Context, obj *model.Room) (string, error) {
	return fmt.Sprintf(roomIDFormat, obj.ID), nil
}

func (r *roomResolver) UserCount(ctx context.Context, obj *model.Room) (int, error) {
	id, err := strconv.Atoi(obj.ID)
	if err != nil {
		return 0, err
	}

	count, err := middleware.GetRoomUserCountLoader(ctx).Load(id)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
