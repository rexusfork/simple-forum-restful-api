package membership

import (
	"context"

	"simple-forum/internals/models/membership"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*membership.UserModel, error) {
	query := "SELECT id, email, username, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = $1 OR username = $2"
	row := r.db.QueryRowContext(ctx, query, email, username)
	var response membership.UserModel

	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
