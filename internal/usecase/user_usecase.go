package usecase

import (
	"context"
	"fmt"
	"github.com/obrunogonzaga/go-template/internal/entity"
)

type UserUseCase interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, id string, input entity.UpdateUserInput) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, page, limit int) ([]entity.User, error)
}

type userUseCase struct {
	userRepo entity.UserRepository
}

func NewUserUseCase(userRepo entity.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(ctx context.Context, user *entity.User) error {
	return uc.userRepo.Create(ctx, user)
}

func (uc *userUseCase) GetByID(ctx context.Context, id string) (*entity.User, error) {
	if id == "" {
		return nil, entity.ErrInvalidInput
	}

	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	if user == nil {
		return nil, entity.ErrUserNotFound
	}

	return user, nil
}

func (uc *userUseCase) List(ctx context.Context, page, limit int) ([]entity.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return uc.userRepo.List(ctx, page, limit)
}

func (uc *userUseCase) Update(ctx context.Context, id string, input entity.UpdateUserInput) error {
	// Verify if user exists
	existing, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return entity.ErrUserNotFound
	}

	return uc.userRepo.Update(ctx, id, input)
}

func (uc *userUseCase) Delete(ctx context.Context, id string) error {
	// Verify if user exists
	existing, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return entity.ErrUserNotFound
	}

	return uc.userRepo.Delete(ctx, id)
}
