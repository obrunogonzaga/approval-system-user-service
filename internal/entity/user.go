package entity

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidRole = errors.New("invalid role")
	ErrInvalidDepartment = errors.New("invalid department")
	ErrInvalidName = errors.New("name must between 3 and 100 characters")
	ErrInvalidEmail = errors.New("invalid email format")
)

type User struct {
	ID         string    	`json:"id"`
	Name       string    	`json:"name" binding:"required,min=3,max=100"`
	Email      string    	`json:"email" binding:"required,email"`
	Role       Role    		`json:"role" binding:"required,oneof=admin developer"`
	Department Department   `json:"department" binding:"required,oneof=developer admin"`
	CreatedAt  time.Time 	`json:"created_at"`
	UpdatedAt  time.Time 	`json:"updated_at"`
}

func NewUser(name, email string, role Role, department Department) (*User, error) {
	user := &User{
		Name:      	name,
		Email:     	email,
		Role: 	   	role,
		Department: department,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if len(u.Name) < 3 || len(u.Name) > 100 {
		return ErrInvalidName
	}
	//TODO: Criar uma função para validar o email (isValidEmail)
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if !u.Role.IsValid() {
		return ErrInvalidRole
	}
	if !u.Department.IsValid() {
		return ErrInvalidDepartment
	}
	
	return nil
}

func (r Role) String() string {
    return string(r)
}

func (d Department) String() string {
    return string(d)
}

func NewRole(r string) (Role, error) {
    role := Role(r)
    if !role.IsValid() {
        return "", fmt.Errorf("invalid role: %s", r)
    }
    return role, nil
}

func NewDepartment(d string) (Department, error) {
    dept := Department(d)
    if !dept.IsValid() {
        return "", fmt.Errorf("invalid department: %s", d)
    }
    return dept, nil
}

type UpdateUserInput struct {
	Name  		string 		`json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Email 		string 		`json:"email,omitempty" binding:"omitempty,email"`
	Role 		*Role 		`json:"role,omitempty"`
	Department 	*Department `json:"department,omitempty"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, input UpdateUserInput) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, page, limit int) ([]User, error)
	Count(ctx context.Context) (int, error)
}
