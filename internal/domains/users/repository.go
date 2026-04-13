package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) DoesUserExist(ctx context.Context, application string, user string) (bool, error) {
	var exists bool

	err := r.DB.QueryRow(ctx,
		"SELECT EXISTS (SELECT 1 FROM applications_users WHERE application=$1 and username=$2)",
		application, user,
	).Scan(&exists)

	return exists, err
}

func (r *Repository) CreateUser(ctx context.Context, application string, user string) error {
	_, err := r.DB.Exec(ctx,
		"insert into applications_users(application, username) values ($1, $2)",
		application, user,
	)

	return err
}

func (r *Repository) RemoveUser(ctx context.Context, application string, user string) error {
	_, err := r.DB.Exec(ctx,
		"delete from applications_users where application=$1 and username=$2",
		application, user,
	)

	return err
}
