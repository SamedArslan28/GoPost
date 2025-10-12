package repository

import (
	"context"
	"database/sql"
	"errors"

	apperrors "github.com/SamedArslan28/gopost/internal/errors"
	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
	INSERT INTO users (username, email, password, created_at)
	VALUES ($1, $2, $3, NOW())
	RETURNING id, created_at
	`

	err := u.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.Id, &user.Created)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, apperrors.ErrEmailConflict
			}
		}
		return nil, err
	}
	return user, nil
}

func (u userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE email = $1`
	row := u.db.QueryRowContext(ctx, query, email)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}

func (u userRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`
	row := u.db.QueryRowContext(ctx, query, id)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}
