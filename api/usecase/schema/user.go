package schema

type UserOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserInput struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// サインイン時の入力
type SignInInput struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// サインアップ時の入力
type SignUpInput struct {
	ID       string   `json:"id" binding:"required"`
	Name     string   `json:"name"`
	Password string   `json:"password" binding:"required"`
	ImageURL []string `json:"image_url" binding:"required"`
}

// サインイン・サインアウト時の出力
type AuthOutput struct {
	Jwt string `json:"jwt"`
}

// 写真の追加時のインプット
type ImageInput struct {
	UserID   string   `json:"user_id" binding:"required"`
	ImageURL []string `json:"image_url" binding:"required"`
}

// type Sastoken struct {
// 	Sas string
// }
