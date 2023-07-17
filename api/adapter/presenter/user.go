package presenter

/*
presenter パッケージは，出力に対するアダプターです．

ここでは，アウトプットポートを実装します(interfaceを満たすようにmethodを追加するということ)
*/

import (
	"log"
	"net/http"

	"megamouth/api/entity/models"
	"megamouth/api/usecase/port"
	"megamouth/api/usecase/schema"
	"megamouth/api/utils/tools"

	"github.com/gin-gonic/gin"
)

type User struct {
	ctx *gin.Context
}

// NewUserOutputPort はUserOutputPortを取得します．
func NewUserOutputPort(ctx *gin.Context) port.UserOutputPort {
	return &User{
		ctx: ctx,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力します．
func (u *User) Render(user *schema.UserOutput) {
	u.ctx.JSON(http.StatusOK, &user)

}

func (u *User) RenderJWTwithUser(user *models.User, img []models.Image) {
	jwt, _ := tools.GenerateJWT(user)
	userOutput := schema.UserOutput{ID: user.ID, Name: user.Name}
	imgOutput := make([]string, len(img))
	for i, v := range img {
		imgOutput[i] = v.URL
	}
	u.ctx.JSON(http.StatusOK, schema.AuthOutputWithUser{Jwt: jwt.Jwt, User: userOutput, ImageURL: imgOutput})
}

func (u *User) RenderJWT(user *models.User) {
	jwt, _ := tools.GenerateJWT(user)
	u.ctx.JSON(http.StatusOK, &jwt)
}

// RenderError はErrorを出力します．
func (u *User) RenderError(err error) {
	u.ctx.JSON(http.StatusInternalServerError, err)
	log.Print(err)
}

func (u *User) RenderIsUsed(message bool) {
	m := schema.MessageOutput{Message: message}
	u.ctx.JSON(http.StatusOK, &m)
}
