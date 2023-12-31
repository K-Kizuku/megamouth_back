package controller

/*
controller パッケージは，入力に対するアダプターです．

ここでは，インプットポートとアウトプットポートを組み立てて，
インプットポートを実行します．
ユースケースレイヤからの戻り値を受け取って出力する必要はなく，
純粋にhttpを受け取り，ユースケースを実行します．
*/

import (
	"megamouth/api/entity/repository"
	"megamouth/api/usecase/port"

	_ "megamouth/api/entity/models"
	_ "megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Post struct {
	OutputFactory func(ctx *gin.Context) port.PostOutputPort
	InputFactory  func(o port.PostOutputPort, p repository.PostRepository) port.PostInputPort
	RepoFactory   func(c *gorm.DB) repository.PostRepository
	Conn          *gorm.DB
}

// @Summary 投稿の投稿
// @Tags post
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param       body body     schema.PostInput false "postの新規作成"
// @Success 200 {object} schema.PostOutput
// @Failure 400 {object} error
// @Router /post/create [post]
func (p *Post) CreatePost(ctx *gin.Context) {
	outputPort := p.OutputFactory(ctx)
	repository := p.RepoFactory(p.Conn)
	inputPort := p.InputFactory(outputPort, repository)
	inputPort.CreatePost(ctx)
}

// @Summary 投稿の全件取得
// @Tags post
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Success 200
// @Failure 400 {object} error
// @Router /post/all [get]
func (p *Post) GetPosts(ctx *gin.Context) {
	outputPort := p.OutputFactory(ctx)
	repository := p.RepoFactory(p.Conn)
	inputPort := p.InputFactory(outputPort, repository)
	inputPort.GetPosts(ctx)
}

// @Summary post_idから投稿を取得
// @Tags post
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param       id   path   int   false "ID"
// @Success 200 {object} schema.PostOutput
// @Failure 400 {object} error
// @Router /post/{id} [get]
func (p *Post) GetPostByID(ctx *gin.Context) {
	outputPort := p.OutputFactory(ctx)
	repository := p.RepoFactory(p.Conn)
	inputPort := p.InputFactory(outputPort, repository)
	inputPort.GetPostByID(ctx)
}

// @Summary 投稿の編集
// @Tags post
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param       body body     schema.PostInput false "postの編集"
// @Param       id   path   int   false  "ID"
// @Success 200 {object} schema.PostOutput
// @Failure 400 {object} error
// @Router /post/{id} [put]
func (p *Post) UpdatePostByID(ctx *gin.Context) {
	outputPort := p.OutputFactory(ctx)
	repository := p.RepoFactory(p.Conn)
	inputPort := p.InputFactory(outputPort, repository)
	inputPort.UpdatePostByID(ctx)
}

// @Summary 投稿の削除
// @Tags post
// @Produce  json
// @securityDefinitions.basic BearerAuth
// @Param       id   path   int   false "ID"
// @Success 200 {object} schema.PostOutput
// @Failure 400 {object} error
// @Router /post/{id} [delete]
func (p *Post) DeletepostByID(ctx *gin.Context) {
	outputPort := p.OutputFactory(ctx)
	repository := p.RepoFactory(p.Conn)
	inputPort := p.InputFactory(outputPort, repository)
	inputPort.DeletepostByID(ctx)
}
