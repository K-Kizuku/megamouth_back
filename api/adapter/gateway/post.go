package gateway

import (
	"megamouth/api/entity/models"
	"megamouth/api/entity/repository"
	"megamouth/api/usecase/schema"
	"megamouth/api/utils/codes"
	"megamouth/api/utils/errors"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostRepository struct {
	conn *gorm.DB
}

func NewPostRepository(conn *gorm.DB) repository.PostRepository {
	return &PostRepository{
		conn: conn,
	}
}

func (p *PostRepository) CreatePost(ctx *gin.Context) (*models.Post, error) {
	conn := p.GetDBConn()
	input := schema.PostInput{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, errors.New(codes.CodeBadRequest, "bad request")
	}
	post := models.Post{Author: input.UserID, Content: input.Content}
	if err := conn.Create(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, errors.New(codes.CodeDatabase, "faild create post")
		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/CreatePost"))
	}

	return &post, nil
}

func (p *PostRepository) GetPosts(ctx *gin.Context) ([]models.Post, error) {
	conn := p.GetDBConn()
	post := []models.Post{}
	if err := conn.Find(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(codes.CodeNotFound, "posts not found")

		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/GetPosts"))
	}
	return post, nil
}

func (p *PostRepository) GetPostByID(ctx *gin.Context) (*models.Post, error) {
	conn := p.GetDBConn()
	post := models.Post{}
	postID := ctx.Param("id")
	if err := conn.First(&post, "id = ?", postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(codes.CodeNotFound, "post not found")
		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/GetPostByID"))
	}
	return &post, nil
}

func (p *PostRepository) UpdatePostByID(ctx *gin.Context) (*models.Post, error) {
	conn := p.GetDBConn()
	input := schema.PostInput{}
	postID := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, errors.New(codes.CodeBadRequest, "bad request")
	}
	post := models.Post{}
	if err := conn.Model(&post).Where("id = ?", postID).Update("content", input.Content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(codes.CodeNotFound, "post not found")
		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/UpdatePostByID"))
	}
	return &post, nil
}

func (p *PostRepository) DeletepostByID(ctx *gin.Context) (*models.Post, error) {
	conn := p.GetDBConn()
	post := models.Post{}
	postID := ctx.Param("id")
	if err := conn.Model(&post).Where("id = ?", postID).Update("deleted_at", time.Now()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(codes.CodeNotFound, "post not found")
		}
		return nil, errors.New(codes.CodeInternal, codes.CodeInternal.DetailString("adapter/gateway/UpdatePostByID"))
	}
	return &post, nil
}

// GetDBConn はconnectionを取得します．
func (p *PostRepository) GetDBConn() *gorm.DB {
	return p.conn
}
