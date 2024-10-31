package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/obrunogonzaga/project/internal/entity"
)

type PaginatedUsers struct {
	Users      []entity.User `json:"users"`
	TotalItems int          `json:"total_items"`
	TotalPages int          `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
	ItemsPerPage int        `json:"items_per_page"`
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, id string, input entity.UpdateUserInput) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, page, limit string) ([]entity.User, error)
}

type userUseCase struct {
	userRepo entity.UserRepository
}

func NewUserUseCase(userRepo entity.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	return uc.userRepo.Create(ctx, user)
}

func (uc *userUseCase) GetUser(ctx context.Context, id string) (*entity.User, error) {
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

func (uc *userUseCase) ListUsers(ctx context.Context, pageStr, limitStr string) (*PaginatedUsers, error) {
	// Valores padrão para paginação
	page := 1
	limit := 10
	
	// Parse página
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil {
			return nil, entity.ErrInvalidInput
		}
		if parsedPage < 1 {
			return nil, entity.ErrInvalidInput
		}
		page = parsedPage
	}
	
	// Parse limite
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			return nil, entity.ErrInvalidInput
		}
		if parsedLimit < 1 || parsedLimit > 100 {
			return nil, entity.ErrInvalidInput
		}
		limit = parsedLimit
	}
	
	// Busca usuarios e total
	users, err := uc.userRepo.List(ctx, page, limit)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}
	
	totalItems, err := uc.userRepo.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("error counting users: %w", err)
	}
	
	// Calcula total de páginas
	totalPages := (totalItems + limit - 1) / limit
	
	return &PaginatedUsers{
		Users:        users,
		TotalItems:   totalItems,
		TotalPages:   totalPages,
		CurrentPage:  page,
		ItemsPerPage: limit,
	}, nil
}

func (uc *userUseCase) UpdateUser(ctx context.Context, id string, input entity.UpdateUserInput) error {
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

func (uc *userUseCase) DeleteUser(ctx context.Context, id string) error {
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