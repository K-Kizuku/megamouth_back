package interactor

import (
	"megamouth/api/entity/repository"
	"megamouth/api/usecase/port"
	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
)

type Post struct {
	OutputPort port.PostOutputPort
	PostRepo   repository.PostRepository
}

// user作成
func (p *Post) CreatePost(ctx *gin.Context) {
	post, err := p.PostRepo.CreatePost(ctx)
	if err != nil {
		p.OutputPort.RenderError(err)
		return
	}
	res := schema.PostOutput{ID: post.ID, Author: post.Author, Content: post.Content, Reply: nil, Reaction: nil, CreatedAt: post.CreatedAt, UpdatedAt: post.UpdatedAt}
	p.OutputPort.Render(&res)
}

// 投稿の全件取得
func (p *Post) GetPosts(ctx *gin.Context) {
	post, err := p.PostRepo.GetPosts(ctx)
	if err != nil {
		p.OutputPort.RenderError(err)
		return
	}
	res := schema.PostOutput{ID: post.ID, Author: post.Author, Content: post.Content, Reply: nil, Reaction: nil, CreatedAt: post.CreatedAt, UpdatedAt: post.UpdatedAt}
	p.OutputPort.Render(&res)
}

// 投稿の1件取得
func (p *Post) GetPostByID(ctx *gin.Context) {
	post, err := p.PostRepo.GetPostByID(ctx)
	if err != nil {
		p.OutputPort.RenderError(err)
		return
	}
	res := schema.PostOutput{ID: post.ID, Author: post.Author, Content: post.Content, Reply: nil, Reaction: nil, CreatedAt: post.CreatedAt, UpdatedAt: post.UpdatedAt}
	p.OutputPort.Render(&res)
}

// 投稿の更新
func (p *Post) UpdatePostByID(ctx *gin.Context) {
	post, err := p.PostRepo.UpdatePostByID(ctx)
	if err != nil {
		p.OutputPort.RenderError(err)
		return
	}
	res := schema.PostOutput{ID: post.ID, Author: post.Author, Content: post.Content, Reply: nil, Reaction: nil, CreatedAt: post.CreatedAt, UpdatedAt: post.UpdatedAt}
	p.OutputPort.Render(&res)
}

// 投稿の削除
func (p *Post) DeletepostByID(ctx *gin.Context) {
	post, err := p.PostRepo.DeletepostByID(ctx)
	if err != nil {
		p.OutputPort.RenderError(err)
		return
	}
	res := schema.PostOutput{ID: post.ID, Author: post.Author, Content: post.Content, Reply: nil, Reaction: nil, CreatedAt: post.CreatedAt, UpdatedAt: post.UpdatedAt}
	p.OutputPort.Render(&res)
}

func NewPostInputPort(outputPort port.PostOutputPort, postRepository repository.PostRepository) port.PostInputPort {
	return &Post{
		OutputPort: outputPort,
		PostRepo:   postRepository,
	}
}
