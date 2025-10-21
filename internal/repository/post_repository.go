package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgconn"
)

type PostRepository interface {
	NewPost(ctx context.Context, authorId int32, title, body string) (*models.Post, error)
	GetAllPostsForUser(ctx context.Context, userId int32) ([]*models.Post, error)
	GetPostById(ctx context.Context, postId int32) (*models.Post, error)
	DeletePost(ctx context.Context, id int32) error
	UpdatePost(ctx context.Context, id int32, title string, body string) (*models.Post, error)
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) NewPost(ctx context.Context, authorId int32, title, body string) (*models.Post, error) {
	query := `
    INSERT INTO posts (title, body, author_id)
    VALUES ($1, $2, $3)
    RETURNING id, title, body, author_id
`

	var post models.Post
	err := p.db.QueryRowContext(ctx, query, title, body, authorId).Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.AuthorId,
	)
	if err != nil {
		var pqErr *pgconn.PgError
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				return nil, fmt.Errorf("a post with the same title already exists: %w", err)
			}
		}
		return nil, fmt.Errorf("failed to insert post: %w", err)
	}

	return &post, nil
}

func (p *postRepository) GetAllPostsForUser(ctx context.Context, userId int32) ([]*models.Post, error) {
	query := `
    SELECT id, title, body, author_id
    FROM posts
    WHERE author_id = $1
`
	rows, err := p.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Error(err)
		}
	}(rows)

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.AuthorId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, &post)
	}

	// Return an empty slice, not an error, if no posts are found
	return posts, nil
}

func (p *postRepository) GetPostById(ctx context.Context, postId int32) (*models.Post, error) {
	query := `
    SELECT id, title, body, author_id
    FROM posts
    WHERE id = $1
`
	var post models.Post
	err := p.db.QueryRowContext(ctx, query, postId).Scan(
		&post.Id,
		&post.Title,
		&post.Body,
		&post.AuthorId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("failed to fetch post: %w", err)
	}

	return &post, nil
}

func (p *postRepository) DeletePost(ctx context.Context, id int32) error {
	var deletedID int32

	query := `DELETE FROM posts WHERE id = $1 RETURNING id`

	err := p.db.QueryRowContext(ctx, query, id).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sql.ErrNoRows
		}
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}

func (p *postRepository) UpdatePost(ctx context.Context, id int32, title string, body string) (*models.Post, error) {
	var updatedPost models.Post

	query := `
       UPDATE posts 
       SET title = $1, body = $2
       WHERE id = $3
       RETURNING id, title, body, author_id`

	err := p.db.QueryRowContext(ctx, query,
		title,
		body,
		id,
	).Scan(
		&updatedPost.Id,
		&updatedPost.Title,
		&updatedPost.Body,
		&updatedPost.AuthorId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return &updatedPost, nil
}
