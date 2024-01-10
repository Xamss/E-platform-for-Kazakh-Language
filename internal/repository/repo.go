package repository

import (
	"context"

	"xamss.diploma/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, u *model.User) error
	GetUser(ctx context.Context, username string) (*model.User, error)
	//Login(ctx context.Context, username, password string) (string, error)
	//UpdateUser(ctx context.Context, u *model.User) error
	//DeleteUser(ctx context.Context, id int64) error

	// CreateActivity(ctx context.Context, a *model.Activity, u model.User) error
	// UpdateActivity(ctx context.Context, a *model.Activity) error
	// DeleteActivity(ctx context.Context, id int64) error
	// GetActivityByID(ctx context.Context, id int64) (*model.Activity, error)
	// GetAllActivities(ctx context.Context) ([]model.Activity, error)
	// GetActivitiesByUserID(ctx context.Context, userID int64) ([]model.Activity, error)
}
