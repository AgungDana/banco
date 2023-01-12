package user

import "context"

type Manager interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*uint, error)
	UpdateUserById(ctx context.Context, req UpdateUserRequest) (*uint, error)
	UpdateUserByEmail(ctx context.Context, req UpdateUserRequest) (*uint, error)
	FindUserByID(ctx context.Context) (*Profile, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	DeleteUser(ctx context.Context) (*uint, error)
}

type Repo interface {
	NewMutation(ctx context.Context) Mutation
	NewQuery() Query
}

type Mutation interface {
	UpdateUserByID(ctx context.Context, user User) (*uint, error)
	UpdateUserByEmail(ctx context.Context, user User) (*uint, error)
	DeletUserByID(ctx context.Context, id uint) error
	CreateUser(ctx context.Context, user User) (*uint, error)

	Commit(ctx context.Context) error
	Cancel(ctx context.Context) error
}

type Query interface {
	FindUserByID(ctx context.Context, id uint) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
