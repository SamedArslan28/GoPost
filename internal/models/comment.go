package models

import "time"

type Comment struct {
	Id        int32     `json:"id"`
	PostID    int32     `json:"postId"`
	UserID    int32     `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
