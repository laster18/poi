// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateRoomInput struct {
	Name string `json:"name"`
}

type Message struct {
	ID            string    `json:"id"`
	UserName      string    `json:"userName"`
	UserAvatarURL string    `json:"userAvatarUrl"`
	Text          string    `json:"text"`
	CreatedAt     time.Time `json:"createdAt"`
}

type Room struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Messages []*Message `json:"messages"`
}

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}
