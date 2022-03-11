package repo

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"teswir-go/internal/entity"
	"teswir-go/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

const (
	sqlUserAdd           = `INSERT INTO tbl_user(username, firstname, lastname, user_role) VALUES($1, $2, $3, $4)`
	sqlUserGetByUsername = `SELECT id, firstname, lastname, user_role, create_ts, update_ts FROM tbl_user WHERE username=$1`
	sqlUserGetByID       = `SELECT username, firstname, lastname, user_role, create_ts, update_ts FROM tbl_user WHERE id=$1`
)

func (u UserRepo) UserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error) {

	row := u.Pool.QueryRow(ctx, sqlUserGetByID, id)
	var user entity.User
	err = row.Scan(&user.Username, &user.Firstname, &user.Lastname, &user.UserRole, &user.CreateTS, &user.UpdateTS)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return
	}

	item = &entity.User{
		ID:        id,
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		UserRole:  user.UserRole,
		CreateTS:  user.CreateTS,
		UpdateTS:  user.UpdateTS,
	}

	return
}

func (u UserRepo) UserGetByUsername(ctx context.Context, username string) (item *entity.User, err error) {

	row := u.Pool.QueryRow(ctx, sqlUserGetByUsername, username)
	var user entity.User
	err = row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.UserRole, &user.CreateTS, &user.UpdateTS)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return
	}

	item = &entity.User{
		ID:        user.ID,
		Username:  username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		UserRole:  user.UserRole,
		CreateTS:  user.CreateTS,
		UpdateTS:  user.UpdateTS,
	}

	return
}

func (u UserRepo) UserAdd(ctx context.Context, r *entity.User) (err error) {

	_, err = u.Pool.Exec(ctx, sqlUserAdd, r.Username, r.Firstname, r.Lastname, r.UserRole)
	return
}
