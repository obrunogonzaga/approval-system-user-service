package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/obrunogonzaga/go-template/internal/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	user.ID = uuid.New().String()

	now := time.Now()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	query := `
		INSERT INTO users (id, name, email, role, department, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		string(user.Role),
		string(user.Department),
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `
		SELECT id, name, email, role, department, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`

	user := &entity.User{}
	var roleStr, deptStr string

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&roleStr,
		&deptStr,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	user.Role = entity.Role(roleStr)
	user.Department = entity.Department(deptStr)

	return user, nil
}

func (r *userRepository) List(ctx context.Context, page, limit int) ([]entity.User, error) {
	offset := (page - 1) * limit

	query := `
        SELECT id, name, email, role, department, created_at, updated_at 
        FROM users 
        ORDER BY created_at DESC 
        LIMIT $1 OFFSET $2
    `

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		var roleStr, deptStr string

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&roleStr,
			&deptStr,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}

		user.Role = entity.Role(roleStr)
		user.Department = entity.Department(deptStr)

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error counting users: %w", err)
	}
	return count, nil
}

func (r *userRepository) Update(ctx context.Context, id string, input entity.UpdateUserInput) error {
	query := `
		UPDATE users 
		SET 
			name = COALESCE($1, name),
			email = COALESCE($2, email),
			role = COALESCE($3, role),
			department = COALESCE($4, department),
			updated_at = $5
		WHERE id = $6
	`

	var roleStr, deptStr *string
	if input.Role != nil {
		s := string(*input.Role)
		roleStr = &s
	
	}
	if input.Department != nil {
		s := string(*input.Department)
		deptStr = &s
	}

	result, err := r.db.ExecContext(
		ctx,
		query,
		input.Name,
		input.Email,
		roleStr,
		deptStr,
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return entity.ErrUserNotFound
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return entity.ErrUserNotFound
	}

	return nil
}
