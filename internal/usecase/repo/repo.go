package repo

import "teswir-go/pkg/postgres"

type Repo struct {
	*postgres.Postgres
}

func NewRepo(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}
