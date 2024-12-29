package generate

const (
	postgresRepositoryTemplate = `package %s

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pxgpool"
	"%s/internal"
)

type I%sRepository interface{}

type %sRepository struct {
	Conn *pgxpool.Pool
}

func New%sRepository(connection *pgxpool.Pool) I%sRepository {
	return &%sRepository{Conn: connection}
}
`

	mysqlRepositoryTemplate = `package %s

import (
	"context"
	"database/sql"
	"%s/internal"
)

type I%sRepository interface{}

type %sRepository struct {
	Conn *sql.DB
}

func New%sRepository(connection *sql.DB) I%sRepository {
	return &%sRepository{Conn: connection}
}
`

	sqlite3RepositoryTemplate = `package %s

import (
	"context"
	"database/sql"
	"%s/internal"
)

type I%sRepository interface{}

type %sRepository struct {
	Conn *sql.DB
}

func New%sRepository(connection *sql.DB) I%sRepository {
	return &%sRepository{Conn: connection}
}
`
)
