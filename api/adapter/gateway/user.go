package gateway

/*
gateway パッケージは，DB操作に対するアダプターです．
*/

import (
	"errors"
	"log"
	"megamouth/api/entity/models"
	"megamouth/api/entity/repository"
	"megamouth/api/usecase/schema"
	"megamouth/api/utils/codes"
	"megamouth/api/utils/tools"

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
			log.Print("user not found")
			return nil, nil
		}
		return nil, errors.New(codes.CodeInternal.DetailString("adapter/gateway/GetUserByID"))
	}
	return &user, nil
}

func (u *UserRepository) CreateUser(ctx *gin.Context) (*models.User, error) {
	conn := u.GetDBConn()
	input := schema.SignUpInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("bad request")
	}
	user := models.User{ID: input.ID, Name: input.Name, PasswordHash: tools.Hash(input.Password), Authentication: &models.Authentication{ID: input.ID, PasswordHash: tools.Hash(input.Password)}}
	if err := conn.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, errors.New("faild create user")
		}
		return nil, errors.New(codes.CodeInternal.DetailString("adapter/gateway/CreateUser"))
	}
	images := make([]models.Image, len(input.ImageURL))
	for i, imageURL := range input.ImageURL {
		images[i] = models.Image{URL: imageURL, Owner: input.ID}
	}

	if err := conn.Create(&images).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, errors.New("failed to create user")
		}
		return nil, errors.New(codes.CodeInternal.DetailString("adapter/gateway/CreateUser"))
	}
	return &user, nil
}

func (u *UserRepository) LoginUser(ctx *gin.Context) (*models.User, error) {
	conn := u.GetDBConn()
	input := schema.SignInInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, errors.New("bad request")
	}
	user := models.User{}
	if err := conn.Where("id = ?", input.ID).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if tools.Hash(input.Password) != user.PasswordHash {
		return nil, errors.New("password invalid")
	}
	return &user, nil
}
func (u *UserRepository) IsUsedName(ctx *gin.Context) (*models.User, error) {
	conn := u.GetDBConn()
	user := models.User{}
	userID := ctx.Param("id")
	if err := conn.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print("user not found")
			return nil, nil
		}
		return nil, errors.New(codes.CodeInternal.DetailString("adapter/gateway/GetUserByID"))
	}
	return &user, nil
}

// GetDBConn はconnectionを取得します．
func (u *UserRepository) GetDBConn() *gorm.DB {
	return u.conn
}
