package entity

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required,min=3,max=100"`
	Email     string    `json:"email" binding:"required,email`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name, email string) *User {
	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type UpdateUserInput struct {
	Name  string `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Email string `json:"email,omitempty" binding:"omitempty,email"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, input UpdateUserInput) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, page, limit int) ([]User, error)
	Count(ctx context.Context) (int, error)
}
