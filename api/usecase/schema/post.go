package schema

import "time"

type PostInput struct {
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostOutput struct {
	ID        uint
	Author    string
	Content   string
	Reply     *PostOutput
	Reaction  []*ReactionOutput
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ReplyInput struct {
	PostID  uint   `json:"post_id" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
