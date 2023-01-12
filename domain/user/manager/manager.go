package manager

import (
	"banco/common/bcrypt"
	"banco/common/config"
	"banco/common/ctxutil"
	"banco/common/jwt"
	"banco/domain/user"
	"banco/domain/user/repo"
	"context"
	"errors"
	"log"
	"time"
)

func NewUserManager(conf config.Config) user.Manager {
	repo, _ := repo.NewRepoUser(conf)
	return &UserManager{
		repo: repo,
	}
}

type UserManager struct {
	repo user.Repo
}

// CreateUser implements user.Manager
func (m *UserManager) CreateUser(ctx context.Context, req user.CreateUserRequest) (*uint, error) {
	mutation := m.repo.NewMutation(ctx)

	newUser := user.NewUser(req)

	pass, err := bcrypt.HashPassword(newUser.Password)
	if err != nil {
		mutation.Cancel(ctx)
		return nil, err
	}
	newUser.Password = pass

	data, err := mutation.CreateUser(ctx, newUser)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}

	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil
}

// DeleteUser implements user.Manager
func (m *UserManager) DeleteUser(ctx context.Context) (*uint, error) {
	mutation := m.repo.NewMutation(ctx)
	id, _ := ctxutil.GetUserIdFromCtx(ctx)
	if id == nil {
		return nil, errors.New("internal server error")
	}
	err := mutation.DeletUserByID(ctx, *id)
	if err != nil {
		mutation.Cancel(ctx)
		return nil, err
	}

	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return id, nil
}

// Login implements user.Manager
func (m *UserManager) Login(ctx context.Context, req user.LoginRequest) (*user.LoginResponse, error) {
	query := m.repo.NewQuery()

	data, err := query.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if ok := bcrypt.CompareHashPassword(req.Password, data.Password); !ok {
		return nil, errors.New("password not match")
	}

	maker := jwt.NewJwtToken("1234567890")
	token, err := maker.GeneratedToken(jwt.Payload{
		UserID:    data.Id,
		Email:     data.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().AddDate(1, 0, 0),
	})

	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Token: token,
		Name:  data.Name,
		Email: data.Email,
	}, nil
}

// FindUserByID implements user.Manager
func (m *UserManager) FindUserByID(ctx context.Context) (*user.Profile, error) {
	query := m.repo.NewQuery()

	id, _ := ctxutil.GetUserIdFromCtx(ctx)
	if id == nil {
		return nil, errors.New("internal server error")
	}

	data, err := query.FindUserByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	profile := user.ToProfile(*data)
	return &profile, nil
}

// UpdateUserByEmail implements user.Manager
func (m *UserManager) UpdateUserByEmail(ctx context.Context, req user.UpdateUserRequest) (*uint, error) {
	mutation := m.repo.NewMutation(ctx)

	id, _ := ctxutil.GetUserIdFromCtx(ctx)
	if id == nil {
		return nil, errors.New("internal server error")
	}

	reqUpdate := user.UpdateUser(req, *id)

	data, err := mutation.UpdateUserByEmail(ctx, reqUpdate)
	if err != nil {
		mutation.Cancel(ctx)
		return nil, err
	}

	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateUserById implements user.Manager
func (m *UserManager) UpdateUserById(ctx context.Context, req user.UpdateUserRequest) (*uint, error) {
	mutation := m.repo.NewMutation(ctx)

	id, _ := ctxutil.GetUserIdFromCtx(ctx)
	if id == nil {
		return nil, errors.New("internal server error")
	}

	reqUpdate := user.UpdateUser(req, *id)

	data, err := mutation.UpdateUserByID(ctx, reqUpdate)
	if err != nil {
		mutation.Cancel(ctx)
		return nil, err
	}

	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil
}
