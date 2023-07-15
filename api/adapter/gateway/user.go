package gateway

/*
gateway パッケージは，DB操作に対するアダプターです．
*/

import (
	"megamouth/api/entity/models"
	"megamouth/api/entity/repository"
	"megamouth/api/utils/codes"
	"megamouth/api/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

// NewUserRepository はUserRepositoryを返します．
func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// GetUserByID はDBからデータを取得します．
func (u *UserRepository) GetUserByID(ctx *gin.Context) (*models.User, error) {
	conn := u.GetDBConn()
	user := models.User{}
	userID := ctx.Param("id")
	if err := conn.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(codes.CodeNotFound, "user not found")

		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/GetUserByID"))
	}
	return &user, nil
}

func (u *UserRepository) CreateUser(ctx *gin.Context) (*models.User, error) {
	conn := u.GetDBConn()
	user := models.User{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		return nil, errors.New(codes.CodeBadRequest, "bad request")
	}

	if err := conn.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, errors.New(codes.CodeDatabase, "faild create user")

		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/CreateUser"))
	}
	return &user, nil
}

// GetDBConn はconnectionを取得します．
func (u *UserRepository) GetDBConn() *gorm.DB {
	return u.conn
}
