package persistence

import (
	"context"
	"fmt"
	"log"

	"users_service/ent"
	"users_service/ent/user"
	"users_service/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgrePersister struct {
	client *ent.Client
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func (persister *PostgrePersister) Init() (err error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=root dbname=test password=pass sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return err
	}
	persister.client = client
	return nil
}

func (persister *PostgrePersister) Deinit() {
	persister.client.Close()
}

func NewPostgrePersister() *PostgrePersister {
	persister := new(PostgrePersister)
	if persister.Init() != nil {
		return nil
	}
	return persister
}

func (persister *PostgrePersister) CreateUser(ctx context.Context, user models.User) (id uuid.UUID, err error) {
	dbUser, err := persister.client.User.
		Create().
		SetID(user.Id).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return
	}
	return dbUser.ID, nil
}

func (persister *PostgrePersister) SearchUsername(ctx context.Context, username string) (bool, error) {
	_, err := persister.client.User.
		Query().
		Where(user.Username(username)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("Failed to check username existence: %w", err)
	}
	return true, nil
}

func (persister *PostgrePersister) SearchUserId(ctx context.Context, id uuid.UUID) (bool, error) {
	_, err := persister.client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("Failed to check user id existence: %w", err)
	}
	return true, nil
}

func (persister *PostgrePersister) CheckUser(ctx context.Context, checkUser models.User) (uuid.UUID, error) {
	user, err := persister.client.User.
		Query().
		Where(user.Username(checkUser.Username)).
		Only(ctx)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Failed to check a user: %w", err)
	}
	if user.Password == checkUser.Password {
		return user.ID, nil
	}
	return uuid.UUID{}, nil
}

func (persister *PostgrePersister) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	dbUsers, err := persister.client.User.
		Query().
		All(ctx)
	if err != nil {
		return
	}
	for _, user := range dbUsers {
		users = append(users, models.User{Id: user.ID, Username: user.Username, Password: user.Password})
	}
	return
}

func (persister *PostgrePersister) CreateOrg(ctx context.Context, org models.Org, owner uuid.UUID) (id uuid.UUID, err error) {
	tx, err := persister.client.Tx(ctx)
	dbOrg, err := tx.Org.
		Create().
		SetID(org.Id).
		SetName(org.Name).
		Save(ctx)
	if err != nil {
		err = rollback(tx, fmt.Errorf("Failed to create a org: %w", err))
		return
	}
	// TODO (Membership endpoints): Create enum for roles
	_, err = tx.Membership.
		Create().
		SetRole("owner").
		SetOrg(dbOrg).
		SetMemberID(owner).
		Save(ctx)
	if err != nil {
		err = rollback(tx, fmt.Errorf("Failed to create a membership: %w", err))
		return
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("Failed to create a org")
		return
	}
	return dbOrg.ID, nil
}

func (persister *PostgrePersister) GetAllOrgs(ctx context.Context) (users []models.Org, err error) {
	dbOrgs, err := persister.client.Org.
		Query().
		All(ctx)
	if err != nil {
		return
	}
	for _, org := range dbOrgs {
		users = append(users, models.Org{Id: org.ID, Name: org.Name})
	}
	return
}
