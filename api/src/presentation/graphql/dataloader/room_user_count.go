package dataloader

import (
	"context"
	"time"

	"github.com/laster18/poi/api/graph/generated"
	"github.com/laster18/poi/api/src/domain/room"
)

func NewRoomUserCountLoader(
	ctx context.Context,
	repo room.Repository,
) *generated.RoomUserCountLoader {
	return generated.NewRoomUserCountLoader(generated.RoomUserCountLoaderConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: 100,
		Fetch: func(roomIDs []int) ([]int, []error) {
			counts, err := repo.CountUserByRoomIDs(ctx, roomIDs)

			if err != nil {
				errs := make([]error, len(roomIDs))
				for i := range make([]int, len(roomIDs)) {
					errs[i] = err
				}
				return nil, errs
			}

			return counts, nil
		},
	})
}
