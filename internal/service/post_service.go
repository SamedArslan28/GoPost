package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	apperrors "github.com/SamedArslan28/gopost/internal/errors"
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return post, nil
}

func (p *PostService) UpdatePost(ctx context.Context, id int32, title string, body string, userID int32) (*models.Post, error) {
	postToUpdate, err := p.repo.GetPostById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	jwtUserIDAString := fmt.Sprintf("%d", userID)
	if postToUpdate.AuthorId != jwtUserIDAString {
		return nil, apperrors.ErrForbidden
	}

	post, err := p.repo.UpdatePost(ctx, id, title, body)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// This could happen if the post was deleted between the check and update
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return post, nil
}

func (p *PostService) DeletePost(ctx context.Context, postID int32, userID int32) error {
	postToDelete, err := p.repo.GetPostById(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.ErrNotFound
		}
		return err
	}

	jwtUserIDAString := fmt.Sprintf("%d", userID)
	if postToDelete.AuthorId != jwtUserIDAString {
		return apperrors.ErrForbidden
	}

	err = p.repo.DeletePost(ctx, postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.ErrNotFound
		}
		return err
	}
	return nil
}
