package usecase

import (
	"OAUTH2/models"
	"context"
	"testing"
)

func Test_usecaseUsers_Create(t *testing.T) {
	usecase := NewUsecaseUser()
	ctx := context.Background()
	err := usecase.Create(ctx, &models.UserParam{
		Username: "john",
		Email:    "john@gmail.com",
		Password: "12345",
		Role:     1,
	})
	if err != nil {
		t.Fatal(err)
	}
}
