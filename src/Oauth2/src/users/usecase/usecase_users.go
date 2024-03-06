package usecase

import (
	"OAUTH2/models"
	"OAUTH2/src/users/repositories"
	"context"

	"github.com/labstack/gommon/log"
)

type (
	UsecaseUsers interface {
		Login(ctx context.Context, param *models.UserParam) (*models.UserResponseLogin, error)
		List(ctx context.Context) ([]models.UserResponse, error)
		Create(ctx context.Context, param *models.UserParam) error
		// Detail()
		// Update()
		// Delete()
	}

	usecaseUsers struct {
		repo repositories.UserRepositories
	}
)

func NewUsecaseUser() UsecaseUsers {
	return &usecaseUsers{
		repo: repositories.NewUserRepositories(),
	}
}

func (u *usecaseUsers) List(ctx context.Context) ([]models.UserResponse, error) {
	datas, _, err := u.repo.List(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var reps []models.UserResponse
	for _, user := range datas {
		reps = append(reps, user.Response())
	}

	return reps, nil
}

func (u *usecaseUsers) Login(ctx context.Context, param *models.UserParam) (*models.UserResponseLogin, error) {
	return nil, nil
}

func (u *usecaseUsers) Create(ctx context.Context, param *models.UserParam) error {
	data, err := param.Validate()
	if err != nil {
		return err
	}
	err = u.repo.Create(ctx, data)
	if err != nil {
		log.Error(err)
		return err
	}
	// kirim email
	// kirim data ke kafka
	return nil
}
