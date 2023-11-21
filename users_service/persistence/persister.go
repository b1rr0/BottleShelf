package persistence

import (
	"context"

	"users_service/models"

	"github.com/google/uuid"
)

type Persister interface {
	CreateUser(ctx context.Context, user models.User) (uuid.UUID, error)
	CheckUser(ctx context.Context, user models.User) (uuid.UUID, error)
	SearchUsername(ctx context.Context, username string) (bool, error)
	SearchUserId(ctx context.Context, id uuid.UUID) (bool, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)

	CreateOrg(ctx context.Context, org models.Org, owner uuid.UUID) (uuid.UUID, error)
	GetAllOrgs(ctx context.Context) ([]models.Org, error)
}
