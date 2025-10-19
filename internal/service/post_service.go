package service

import (
	"context"

	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/SamedArslan28/gopost/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (p *PostService) CreatePost(ctx context.Context, post models.Post, authorID int32) (*models.Post, error) {
	newPost, err := p.repo.NewPost(ctx, authorID, post.Title, post.Body)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (p *PostService) GetAllPostForUser(ctx context.Context, authorID int32) ([]*models.Post, error) {
	posts, err := p.repo.GetAllPostsForUser(ctx, authorID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) GetPostById(c context.Context, postID int32) (*models.Post, error) {
	post, err := p.repo.GetPostById(c, postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}
