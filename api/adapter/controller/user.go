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
// @Success 200 {object} error
// @Failure 400 {object} error
// @Router /api/v1/user/:id [get]
func (u *User) GetUserByID(ctx *gin.Context) {

	outputPort := u.OutputFactory(ctx)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx)
}
