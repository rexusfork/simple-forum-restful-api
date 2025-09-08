package membership

import (
	"context"
	"database/sql"
	"log"

	"simple-forum/internals/models/membership"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*membership.UserModel, error) {
	query := "SELECT id, email, username, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ?"
	row := r.db.QueryRowContext(ctx, query, email, username)
	var response membership.UserModel

	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, user membership.UserModel) error {
	query := "INSERT INTO users (email, username, password, created_at, updated_at, created_by, updated_by) VALUES (?,?,?,?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		log.Printf("Error saat membuat user: %v\n", err)
		return err
	}
	return nil
}
