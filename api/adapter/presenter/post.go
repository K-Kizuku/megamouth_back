package presenter

/*
presenter パッケージは，出力に対するアダプターです．

ここでは，アウトプットポートを実装します(interfaceを満たすようにmethodを追加するということ)
*/

import (
	"net/http"

	"megamouth/api/usecase/port"
	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ctx *gin.Context
}

// NewUserOutputPort はUserOutputPortを取得します．
func NewPostOutputPort(ctx *gin.Context) port.PostOutputPort {
	return &Post{
		ctx: ctx,
	}
}

// usecase.UserOutputPortを実装している
// Render はNameを出力します．
func (p *Post) Render(post *schema.PostOutput) {
	p.ctx.JSON(http.StatusOK, &post)

}

// RenderError はErrorを出力します．
func (p *Post) RenderError(err error) {
	p.ctx.JSON(http.StatusInternalServerError, err)

}
