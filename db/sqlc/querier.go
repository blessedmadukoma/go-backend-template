// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	// WHERE email = $1
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)