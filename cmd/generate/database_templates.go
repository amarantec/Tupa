package generate

const (
	postgresDatabasePackageContent = `package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func OpenConnection(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	DB, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(ctx); err != nil {
        DB.Close()
        return nil, fmt.Errorf("database connection ping failed: %w", err)
    }	

	return DB, nil
}`

	mysqlDatabasePackageContent = `package db

import (
	"context"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func OpenConnection(ctx context.Context, connectionString string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	// Verifica se a conexão é válida
	if err = DB.PingContext(ctx); err != nil {
		DB.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return DB, nil
}`

	sqlite3DatabasePackageContent = `package db

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenConnection() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to dabase.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}`
)
