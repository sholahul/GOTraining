package repositories

import (
	"OAUTH2/models"
	"context"
	"testing"

	"github.com/google/uuid"
)

func Test_UserRepositoies_Create(t *testing.T) {
	repo := NewUserRepositories()
	ctx := context.Background()
	err := repo.Create(ctx, &models.User{
		Username: "Siapa",
		Email:    "siapa@gmail.com",
		Password: "12345",
		Role:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_UserRepositoies_Update(t *testing.T) {
	repo := NewUserRepositories()
	ctx := context.Background()

	id := uuid.MustParse("a34e5fe2-cbb1-4209-9586-f6b38460f076")

	data, err := repo.Detail(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	data.Username = "Siapa Aku"

	err = repo.Update(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_UserRepositoies_Delete(t *testing.T) {
	repo := NewUserRepositories()
	ctx := context.Background()
	id := uuid.MustParse("a34e5fe2-cbb1-4209-9586-f6b38460f076")

	err := repo.Delete(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_UserRepositoies_List(t *testing.T) {
	repo := NewUserRepositories()
	ctx := context.Background()
	datas, total, err := repo.List(ctx)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(datas)
	t.Log(total)
}

func Test_UserRepositoies_Detail(t *testing.T) {
	repo := NewUserRepositories()
	ctx := context.Background()

	id := uuid.MustParse("d3926e55-460a-466b-b2a4-e53fe825823f")

	data, err := repo.Detail(ctx, id)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}
