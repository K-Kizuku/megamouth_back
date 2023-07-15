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

	_ "megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type User struct {
	OutputFactory func(ctx *gin.Context) port.UserOutputPort
	// -> presenter.NewUserOutputPort
	InputFactory func(o port.UserOutputPort, u repository.UserRepository) port.UserInputPort
	// -> interactor.NewUserInputPort
	RepoFactory func(c *gorm.DB) repository.UserRepository
	// -> gateway.NewUserRepository
	Conn *gorm.DB
}

// GetUserByID は，httpを受け取り，portを組み立てて，inputPort.GetUserByIDを呼び出します．
// @Summary userIDからUserを返す
// @Tags user
// @Produce  json
// @Param       id   path     string      false "ID"
// @Success 200 {object} schema.UserOutput
// @Failure 400 {object} error
// @Router /user/:id [get]
func (u *User) GetUserByID(ctx *gin.Context) {
	outputPort := u.OutputFactory(ctx)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx)
}

// @Summary userの新規作成
// @Tags user
// @Produce  json
// @Param       body body     schema.SignUpInput false "userの新規作成"
// @Success 200 {object} schema.AuthOutput
// @Failure 400 {object} error
// @Router /user/create [post]
func (u *User) CreateUser(ctx *gin.Context) {
	outputPort := u.OutputFactory(ctx)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.CreateUser(ctx)
}

// @Summary サインイン
// @Tags user
// @Produce  json
// @Param       body body     schema.SignInInput false "サインイン"
// @Success 200 {object} schema.AuthOutput
// @Failure 400 {object} error
// @Router /user/login [post]
func (u *User) LoginUser(ctx *gin.Context) {
	outputPort := u.OutputFactory(ctx)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.LoginUser(ctx)
}
