package services

import (
	"context"
	"errors"

	"github.com/C0deNe0/blog-go/domain"
	"github.com/C0deNe0/blog-go/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, post domain.Post) error
	GetPostById(ctx context.Context, id string) (*domain.Post, error)
	ListPosts(ctx context.Context) ([]domain.Post, error)
	UpdatePost(ctx context.Context, post domain.Post) error
	DeletePost(ctx context.Context, id string) error
}

type postService struct {
	postRepo repository.PostRepository
}

// CreatePost implements PostService.
func (p *postService) CreatePost(ctx context.Context, post domain.Post) error {
	if post.Title == "" || post.Content == "" {
		return errors.New("title and content cannot be empty")
	}
	return p.postRepo.Create(ctx, post)
}

// DeletePost implements PostService.
func (p *postService) DeletePost(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("post Id is required for deletion")
	}
	return p.postRepo.Delete(ctx, id)
}

// GetPostById implements PostService.
func (p *postService) GetPostById(ctx context.Context, id string) (*domain.Post, error) {
	return p.postRepo.GetById(ctx, id)
}

// ListPosts implements PostService.
func (p *postService) ListPosts(ctx context.Context) ([]domain.Post, error) {
	return p.postRepo.List(ctx)
}

// UpdatePost implements PostService.
func (p *postService) UpdatePost(ctx context.Context, post domain.Post) error {
	if post.Id == "" {
		return errors.New("post Id is required to update")
	}
	return p.postRepo.Update(ctx, post)
}

func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{postRepo: postRepo}
}
