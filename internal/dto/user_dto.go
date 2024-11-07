package dto

import "github.com/obrunogonzaga/go-template/internal/entity"

type CreateUserInput struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	Email      string `json:"email" binding:"required,email"`
	Role       string `json:"role" binding:"required"`
	Department string `json:"department" binding:"required"`
}

func (i *CreateUserInput) ToEntity() (*entity.User, error) {
	role, err := entity.ParseRole(i.Role)
	if err != nil {
		return nil, err
	}

	department, err := entity.ParseDepartment(i.Department)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(i.Name, i.Email, role, department)
}
