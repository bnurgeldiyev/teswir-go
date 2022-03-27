package repo

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"teswir-go/internal/entity"
)

const (
	sqlUserAdd           = `INSERT INTO tbl_user(username, firstname, lastname, user_role) VALUES($1, $2, $3, $4)`
	sqlUserGetByUsername = `SELECT id, firstname, lastname, user_role, create_ts, update_ts FROM tbl_user WHERE username=$1`
	sqlUserGetByID       = `SELECT username, firstname, lastname, user_role, create_ts, update_ts FROM tbl_user WHERE id=$1`
	sqlUserList          = `SELECT id, username, firstname, lastname, user_role, create_ts, update_ts FROM tbl_user ORDER BY create_ts`
)

func (u *Repo) RepoUserGetByID(ctx context.Context, id uuid.UUID) (item *entity.User, err error) {

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

func (u *Repo) RepoUserGetByUsername(ctx context.Context, username string) (item *entity.User, err error) {

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

func (u *Repo) RepoUserAdd(ctx context.Context, r *entity.User) (err error) {

	_, err = u.Pool.Exec(ctx, sqlUserAdd, r.Username, r.Firstname, r.Lastname, r.UserRole)
	return
}

func (u *Repo) RepoUserList(ctx context.Context) (item []*entity.User, err error) {

	rows, err1 := u.Pool.Query(ctx, sqlUserList)
	if err1 != nil {
		err = err1
	}

	users := make([]*entity.User, 0)
	for rows.Next() {
		user := new(entity.User)
		err = rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.UserRole, &user.CreateTS, &user.UpdateTS)
		if err != nil {
			return
		}

		users = append(users, user)
	}

	item = users

	return
}
