package schema

// リアクションの登録時・削除時の入力
type ReactionInput struct {
	ReactionID uint   `json:"reaction_id" binding:"required"`
	PostID     uint   `json:"post_id" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}

// リアクションの出力(他のスキーマに埋め込む)
type ReactionOutput struct {
	ReactionID uint        `json:"reaction_id"`
	User       *UserOutput `json:"user"`
}
