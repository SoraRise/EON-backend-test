package storage

import (
	"backend-test/internal/adapter"
	"backend-test/pkg/postgres"
	"context"
	"log/slog"
)

type repository struct {
	pg_client postgres.Pg_Client
	log       *slog.Logger
}

func NewRepository(client postgres.Pg_Client, logger *slog.Logger) adapter.User_adapter {
	return &repository{
		pg_client: client,
		log:       logger,
	}
}

func (r *repository) ParseXML(ctx context.Context) error {
	query := `insert into Apartments (Id, S_total, S_living, S_kitchen, Height, Price, Floor)`
	panic(query)
	return nil
}

// func (r *repository) Create(ctx context.Context, user *domain.Create_user_DTO) (err error, id int) {
// 	var ID int

// 	query := `insert into User (account_name,account_password,email,display_name,image_path,account_status) values($1, $2, $3, $4, $5);`

// 	if err = r.pg_client.QueryRow(ctx, query, user.Account_name, user.Account_password, user.Email, user.Display_name, user.Image_path, user.Account_status).Scan(ID); err != nil {
// 		var pgErr *pgconn.PgError
// 		if errors.As(err, &pgErr) {
// 			pgErr = err.(*pgconn.PgError)
// 			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
// 			r.log.Error("Error", newErr)
// 			return newErr, -1
// 		}
// 		return err, ID
// 	}

// 	return nil, ID
// }