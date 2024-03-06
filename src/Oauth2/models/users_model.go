package models

import (
	"OAUTH2/pkg/hash"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	User struct {
		bun.BaseModel `bun:"table:users,alias:u"`
		ID            uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
		Username      string
		Email         string
		Password      string
		Role          int64     `bun:"default:2"`
		CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
		UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	}

	UserParam struct {
		ID       string `param:"id" query:"id" form:"id" json:"id" xml:"id"`
		Username string `param:"username" query:"username" form:"username" json:"username" xml:"username"`
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
		Role     int64  `param:"role" query:"role" form:"role" json:"role" xml:"role"`
	}

	UserResponse struct {
		ID        string    `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Role      int64     `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserResponseLogin struct {
		User  *UserResponse `json:"user"`
		Token string        `json:"token"`
	}
)

func (data *User) Response() UserResponse {
	return UserResponse{
		ID:        data.ID.String(),
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func (data *UserParam) Validate() (*User, error) {
	// generate password

	if data.Username == "" {
		return nil, errors.New("username di perlukan")
	}
	if data.Email == "" {
		return nil, errors.New("email di perlukan")
	}
	if data.Password == "" {
		return nil, errors.New("password di perlukan")
	}

	newPwd := hash.NewSHA256(data.Password)
	if data.ID == "" {
		data.ID = uuid.NewString()
	}

	return &User{
		ID:        uuid.MustParse(data.ID),
		Username:  data.Username,
		Email:     data.Email,
		Password:  newPwd,
		Role:      data.Role,
		CreatedAt: time.Now(),
	}, nil
}
