package subscriber

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/laster18/poi/api/graph/model"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/infra/redis"
)

type RoomUserSubscriber struct {
	client *redis.Client
	Mutex  sync.Mutex
	// channels map[roomId]map[userId]chan ...
	chs map[int]map[string]chan model.RoomUserEvent
}

func NewRoomUserSubscriber(ctx context.Context, client *redis.Client) *RoomUserSubscriber {
	subscriber := &RoomUserSubscriber{
		client: client,
		Mutex:  sync.Mutex{},
		chs:    make(map[int]map[string]chan model.RoomUserEvent),
	}
	go subscriber.start(ctx)
	return subscriber
}

func (s *RoomUserSubscriber) start(ctx context.Context) {
	// subscribeCh "roomUser:*"
	subscribeCh := fmt.Sprintf("%s:%s:%s", redis.KeySpace, RoomUserChannel, "*")
	pubsub := s.client.PSubscribe(ctx, subscribeCh)
	defer pubsub.Close()

	for {
		msg := <-pubsub.Channel()

		// debug log
		log.Printf("subscribe roomUser, channel: %s, payload: %s\n\n", msg.Channel, msg.Payload)

		ch := removeKeyspacePrefix(msg.Channel)
		roomID, userUID, err := destructRoomUserKey(ch)
		if err != nil {
			log.Println("getted invalid channel key from redis")
			continue
		}

		switch msg.Payload {
		case redis.EventSet:
			roomUserJSON, err := s.client.Get(ctx, ch).Result()
			if err != nil {
				log.Println("failed to get from redis, err:", err)
				continue
			}

			var roomUser domain.RoomUser
			if err := json.Unmarshal([]byte(roomUserJSON), &roomUser); err != nil {
				log.Println("getted unexpected json data struct from redis")
				continue
			}

			d, err := s.makeDataFromSet(&roomUser, roomID, userUID)
			if err != nil {
				log.Println(err)
			}
			s.deliver(roomID, d)
		case redis.EventDel:
			fallthrough
		case redis.EventExpired:
			data := &model.Exited{
				UserID: makeRoomUserID(userUID),
			}
			s.deliver(roomID, data)
		default:
			fmt.Println("received unknown event:", msg.Payload)
		}
	}
}

func (s *RoomUserSubscriber) deliver(roomID int, data model.RoomUserEvent) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	chs, ok := s.chs[roomID]
	if !ok {
		return
	}

	for _, ch := range chs {
		ch <- data
	}
}

func (s *RoomUserSubscriber) Subscribe(ctx context.Context, roomID int, userUID string) <-chan model.RoomUserEvent {
	createdCh := make(chan model.RoomUserEvent)

	go func() {
		<-ctx.Done()
	}()

	return createdCh
}

func (s *RoomUserSubscriber) AddCh(ch chan model.RoomUserEvent, roomID int, userUID string) {
	s.Mutex.Lock()
	userChannels, ok := s.chs[roomID]
	if !ok {
		userChannels = make(map[string]chan model.RoomUserEvent)
		s.chs[roomID] = userChannels
	}
	userChannels[userUID] = ch
	s.Mutex.Unlock()
}

func (s *RoomUserSubscriber) makeDataFromSet(
	ru *domain.RoomUser,
	roomID int,
	userUID string,
) (model.RoomUserEvent, error) {
	switch ru.LastEvent {
	case domain.JoinEvent:
		return &model.Joined{
			UserID:    makeRoomUserID(userUID),
			Name:      ru.Name,
			AvatarURL: ru.AvatarURL,
			X:         ru.X,
			Y:         ru.Y,
		}, nil
	case domain.MoveEvent:
		return &model.Moved{
			UserID: makeRoomUserID(userUID),
			X:      ru.X,
			Y:      ru.Y,
		}, nil
	case domain.MessageEvent:
		if ru.LastMessage == nil {
			return nil, errors.New("not found roomUser.LastMessage")
		}

		return &model.SendedMassage{
			UserID: makeRoomUserID(userUID),
			Message: &model.Message{
				ID:            makeMessageID(ru.LastMessage.ID),
				UserID:        makeUserID(ru.LastMessage.UserUID),
				UserName:      ru.LastMessage.UserName,
				UserAvatarURL: ru.LastMessage.UserAvatarURL,
				Body:          ru.LastMessage.Body,
				CreatedAt:     ru.LastMessage.CreatedAt,
			},
		}, nil
	default:
		return nil, errors.New("getted unknown roomUser event")
	}
}

// TODO: resolverで使っているものと共通化する
func makeRoomUserID(userUID string) string {
	return fmt.Sprintf("RoomUser:%s", userUID)
}

// TODO: resolverで使っているものと共通化する
func makeMessageID(id int) string {
	return fmt.Sprintf("Messaage:%d", id)
}

// TODO: resolverで使っているものと共通化する
func makeUserID(userUID string) string {
	return fmt.Sprintf("User:%s", userUID)
}
