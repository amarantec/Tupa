package utils

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/amarantec/tupa/constants"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func LoadDBConfig() (string, string, error) {
	envFile, err := FindEnvFile()
	if err != nil {
		return constants.EMPTY_STRING, constants.EMPTY_STRING, err
	}

	err = godotenv.Load(envFile)
	if err != nil {
		return constants.EMPTY_STRING, constants.EMPTY_STRING, fmt.Errorf("error loading env file")
	}

	var dbHost, dbPort, dbUser, dbPassword, dbName string

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "postgres" {
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbUser = os.Getenv("POSTGRES_USER")
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
		dbName = os.Getenv("POSTGRES_DB")

		if err := validateEnvVars(dbDriver, dbHost, dbPort, dbUser, dbPassword, dbName); err != nil {
			return constants.EMPTY_STRING, constants.EMPTY_STRING, err
		}
	} else if dbDriver == "mysql" {
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbUser = os.Getenv("MYSQL_USER")
		dbPassword = os.Getenv("MYSQL_PASSWORD")
		dbName = os.Getenv("MYSQL_DATABASE")
		if err := validateEnvVars(dbDriver, dbHost, dbPort, dbUser, dbPassword, dbName); err != nil {
			return constants.EMPTY_STRING, constants.EMPTY_STRING, err
		}
	} else if dbDriver == "sqlite3" {
		dbName = os.Getenv("DB_NAME")
		if err := validateEnvVars(dbDriver, dbName); err != nil {
			return constants.EMPTY_STRING, constants.EMPTY_STRING, err
		}
	}

	var connectionString string
	switch dbDriver {
	case "postgres":
		connectionString = fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, dbHost, dbPort, dbUser, dbPassword, dbName)
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	case "sqlite3":
		connectionString = dbName
	default:
		return constants.EMPTY_STRING, constants.EMPTY_STRING, fmt.Errorf("unsupported database driver: %s", dbDriver)
	}

	return dbDriver, connectionString, nil
}

func GetDbConnection() (DBConnection, error) {
	dbDriver, connectionString, err := LoadDBConfig()
	if err != nil {
		return nil, err
	}

	switch dbDriver {
	case "sqlite3":
		db, err := sql.Open(dbDriver, connectionString)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
		}

		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("failed to ping SQLite database: %w", err)
		}

		return &SQLDB{DB: db}, nil
	case "mysql":
		db, err := sql.Open(dbDriver, connectionString)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
		}

		// Teste de conexão não é tão relevante para SQLite, mas podemos verificar se o arquivo é válido.
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("failed to ping MySQL database: %w", err)
		}

		return &SQLDB{DB: db}, nil
	case "postgres":
		// Postgres usando pgxpool.
		cfg, err := pgxpool.ParseConfig(connectionString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Postgres connection config: %w", err)
		}

		pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to Postgres database: %w", err)
		}

		return &PgxPool{Pool: pool}, nil

	default:
		return nil, fmt.Errorf("unsupported database driver: %s", dbDriver)
	}
}

type DBConnection interface {
	Exec(context.Context, string, ...interface{}) (int64, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (*sql.Rows, error)
	Close() error
}

// wrapper to *sql.DB (mysql, sqlite3)
type SQLDB struct {
	DB *sql.DB
}

func (s *SQLDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return nil
}

func (s *SQLDB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.QueryContext(ctx, query, args...)
}

func (s *SQLDB) Exec(ctx context.Context, query string, args ...interface{}) (int64, error) {
	result, err := s.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *SQLDB) Close() error {
	return s.DB.Close()
}

// Wrapper para *pgxpool.Pool
type PgxPool struct {
	Pool *pgxpool.Pool
}

type PgxRow struct {
	row pgx.Row
}

func (r *PgxRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

func (p *PgxPool) Exec(ctx context.Context, query string, args ...interface{}) (int64, error) {
	result, err := p.Pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func (p *PgxPool) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Criar um wrapper para converter pgx.Rows em *sql.Rows, se necessário.
	return nil, fmt.Errorf("conversion from pgx.Rows to *sql.Rows not implemented")
}

func (p *PgxPool) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	row := p.Pool.QueryRow(ctx, query, args...)
	return &PgxRow{row: row}
}

func (p *PgxPool) Close() error {
	p.Pool.Close()
	return nil
}

func validateEnvVars(vars ...string) error {
	for _, v := range vars {
		if v == constants.EMPTY_STRING {
			return fmt.Errorf("environment variable %s is empty", v)
		}
	}
	return nil
}
