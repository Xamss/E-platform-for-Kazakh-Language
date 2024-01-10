package usecase

import (
	"context"

	"xamss.diploma/internal/model"
)

type Service interface {
	CreateUser(ctx context.Context, u *model.User) error
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(token string) (int64, error)

	//UpdateUser(ctx context.Context, u *entity.User) error
	//DeleteUser(ctx context.Context, id int64) error

	//VerifyToken(token string) error
	//

	// CreateActivity(ctx context.Context, a *model.Activity) error
	// UpdateActivity(ctx context.Context, a *model.Activity) error
	// DeleteActivity(ctx context.Context, id int64) error
	// GetArticleByID(ctx context.Context, id int64) (*model.Activity, error)
}
