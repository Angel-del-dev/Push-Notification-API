package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"notificationapi.com/internal/domains/auth/dtos"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) GetApplicationByDomain(ctx context.Context, application string, key string) (dtos.ApplicationType, error) {
	query := `
		SELECT application, key, password
		FROM applications_keys
		WHERE application = $1
			and key = $2
	`

	var app dtos.ApplicationType

	err := r.DB.QueryRow(ctx, query, application, key).
		Scan(&app.Application, &app.Key, &app.Password)

	return app, err
}
