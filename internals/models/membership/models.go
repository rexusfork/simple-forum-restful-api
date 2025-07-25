package membership

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Email     string    `db:"email"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}
)
