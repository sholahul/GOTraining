package repositories

import (
	"OAUTH2/models"
	"OAUTH2/pkg/db"
	"context"
	"sync"

	"github.com/google/uuid"
)

type (
	UserRepositories interface {
		List(ctx context.Context) ([]models.User, int, error)
		Detail(ctx context.Context, id uuid.UUID) (*models.User, error)
		Create(ctx context.Context, data *models.User) error
		Update(ctx context.Context, data *models.User) error
		Delete(ctx context.Context, id uuid.UUID) error
	}
	userRepositories struct{}
)

var (
	lock = sync.Mutex{}
)

func NewUserRepositories() UserRepositories {
	return &userRepositories{}
}

func (r *userRepositories) List(ctx context.Context) ([]models.User, int, error) {
	lock.Lock()
	defer lock.Unlock()
	var datas = []models.User{}
	total, err := db.GetDB().NewSelect().
		Model(&datas).
		Limit(10).
		ScanAndCount(ctx)

	if err != nil {
		return nil, 0, err
	}
	return datas, total, err
}

func (r *userRepositories) Detail(ctx context.Context, id uuid.UUID) (*models.User, error) {
	lock.Lock()
	defer lock.Unlock()
	var data = &models.User{ID: id}
	err := db.GetDB().NewSelect().Model(data).WherePK().Scan(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *userRepositories) Create(ctx context.Context, data *models.User) error {
	lock.Lock()
	defer lock.Unlock()
	_, err := db.GetDB().NewInsert().Model(data).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (r *userRepositories) Update(ctx context.Context, data *models.User) error {
	lock.Lock()
	defer lock.Unlock()
	_, err := db.GetDB().NewUpdate().Model(data).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (r *userRepositories) Delete(ctx context.Context, id uuid.UUID) error {
	lock.Lock()
	defer lock.Unlock()
	var data = &models.User{ID: id}
	_, err := db.GetDB().NewDelete().Model(data).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
