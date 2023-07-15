package presenter

/*
presenter パッケージは，出力に対するアダプターです．

ここでは，アウトプットポートを実装します(interfaceを満たすようにmethodを追加するということ)
*/

import (
	"log"
	"net/http"

	"megamouth/api/usecase/port"
	"megamouth/api/usecase/schema"

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

// RenderError はErrorを出力します．
func (u *User) RenderError(err error) {
	u.ctx.JSON(http.StatusInternalServerError, err)
	log.Fatal(err)

}
