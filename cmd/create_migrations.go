package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/amarantec/tupa/cmd/utils"
	"github.com/amarantec/tupa/constants"
)

func CreateMigrations(name, sql string) error {
	migrationsPath, err := ensureMigrationsDir()
	if err != nil {
		return err
	}

	driver, err := utils.GetDBDriverFromEnv()
	if err != nil {
		return err
	}
	var transactionSQL string

	switch driver {

	case "postgres":
		transactionSQL = `BEGIN;`
	case "mysql":
		transactionSQL = `START TRANSACTION;`
	case "sqlite3":
		transactionSQL = `BEGIN TRANSACTION;`
	default:
		return fmt.Errorf("unsupported database driver: %s", driver)
	}

	filename := fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102150405"), name)
	migrationsFilePath := filepath.Join(migrationsPath, filename)
	finalSQL := fmt.Sprintf(
		`-- Migration: %s
-- Generated on: %s

%s

%s

COMMIT;`, name, time.Now().Format(time.RFC3339), transactionSQL, sql)
	return writeFile(migrationsFilePath, finalSQL)
}

func GenerateSQLFromStruct(args []string, dbDriver string) (string, error) {
	if len(args) < 2 {
		return constants.EMPTY_STRING, fmt.Errorf("usage: tupa generate <ModelName> <Field:Type> <Field:Type>")
	}

	typeMapping, err := getTypeMapping(dbDriver)
	if err != nil {
		return constants.EMPTY_STRING, err
	}

	return buildSQL(args, typeMapping)
}

func ApplyMigrations() error {
	db, err := utils.GetDbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	migrationsPath, err := ensureMigrationsDir()
	if err != nil {
		return err
	}

	migrationsTable, err := getMigrationsTable()
	if err != nil {
		return err
	}

	if err := createMigrationsTable(db, migrationsTable); err != nil {
		return err
	}

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			if err := applySingleMigration(db, file.Name(), migrationsPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func ensureMigrationsDir() (string, error) {
	internalPath, err := utils.FindProjectInternal()
	if err != nil {
		return constants.EMPTY_STRING, fmt.Errorf("failed to locate internal directory: %w", err)
	}

	migrationsPath := filepath.Join(internalPath, "migrations")
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		if err := os.Mkdir(migrationsPath, os.ModePerm); err != nil {
			return constants.EMPTY_STRING, fmt.Errorf("failed to create migrations directory: %w", err)
		}
	}

	return migrationsPath, nil
}

func getTypeMapping(dbDriver string) (map[string]string, error) {
	switch dbDriver {
	case "postgres":
		return postgresType, nil
	case "mysql":
		return mysqlTypes, nil
	case "sqlite3":
		return sqlite3Types, nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", dbDriver)
	}
}

func buildSQL(args []string, typeMapping map[string]string) (string, error) {
	tableName := args[0]
	fields := args[1:]

	var sql strings.Builder
	sql.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", tableName))

	for i, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			return constants.EMPTY_STRING, fmt.Errorf("invalid field format: %s, expected <FieldName>:<Type>", field)
		}

		fieldName := parts[0]
		fieldType := parts[1]

		sqlType, ok := typeMapping[fieldType]
		if !ok {
			return constants.EMPTY_STRING, fmt.Errorf("unsupported field type: %s", fieldType)
		}

		sql.WriteString(fmt.Sprintf("    %s %s", fieldName, sqlType))
		if i < len(fields)-1 {
			sql.WriteString(",\n")
		} else {
			sql.WriteString("\n")
		}
	}
	sql.WriteString(");")

	upSQL := sql.String()
	downSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)

	return fmt.Sprintf("-- Up Migration\n%s\n\n-- Down Migration\n%s", upSQL, downSQL), nil
}

func getMigrationsTable() (string, error) {
	driver, _ := utils.GetDBDriverFromEnv()
	switch driver {
	case "postgres":
		return migrationsTablePostgres, nil
	case "mysql":
		return migrationsTableMysql, nil
	case "sqlite3":
		return migrationsTableSqlite3, nil
	default:
		return constants.EMPTY_STRING, fmt.Errorf("unsupported database driver: %s", driver)
	}
}

func createMigrationsTable(db utils.DBConnection, migrationsTable string) error {
	_, err := db.Exec(context.Background(), migrationsTable)
	return err
}

func applySingleMigration(db utils.DBConnection, fileName, migrationsPath string) error {
	var count int
	var query string

	driver, _ := utils.GetDBDriverFromEnv()
	if driver == "postgres" {
		query = "SELECT COUNT(*) FROM migrations WHERE migration_name = $1"
	} else {
		query = "SELECT COUNT(*) FROM migrations WHERE migration_name = ?"
	}

	if driver == "postgres" {
		err := db.QueryRow(context.Background(), query, fileName).Scan(&count)
		if err != nil {
			return err
		}
	} else {
		rows, err := db.Query(context.Background(), query, fileName)
		if err != nil {
			return err
		}
		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&count)
			if err != nil {
				return err
			}
		}
	}

	if count > 0 {
		log.Printf("Migration %s already applied, skipping", fileName)
		return nil
	}

	migrationsFilePath := filepath.Join(migrationsPath, fileName)
	sqlStatements, err := os.ReadFile(migrationsFilePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	upSQL, _, err := extractUpAndDownSQL(string(sqlStatements))
	if err != nil {
		return err
	}

	fmt.Println(upSQL)

	// Executa o SQL diretamente
	_, err = db.Exec(context.Background(), upSQL)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to apply migration %s: %w", fileName, err)
	}

	var insertQuery string
	if driver == "postgres" {
		insertQuery = "INSERT INTO migrations (migration_name) VALUES ($1);"
	} else {
		insertQuery = "INSERT INTO migrations (migration_name) VALUES (?);"
	}
	// Registra a migração no banco de dados
	_, err = db.Exec(context.Background(), insertQuery, fileName)
	if err != nil {
		return fmt.Errorf("failed to log migration %s: %w", fileName, err)
	}

	log.Printf("Migration %s applied and logged successfully", fileName)
	return nil
}

func writeFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	return err
}

func extractUpAndDownSQL(sqlStatements string) (string, string, error) {
	parts := strings.Split(sqlStatements, "-- Down Migration")
	if len(parts) < 2 {
		return constants.EMPTY_STRING, constants.EMPTY_STRING, fmt.Errorf("invalid migration format, missing Down Migration")
	}

	upSQL := parts[0]
	downSQL := parts[1]

	upSQL = strings.TrimSpace(upSQL)
	downSQL = strings.TrimSpace(downSQL)

	upSQL = fmt.Sprintf("%s\n\nCOMMIT;", upSQL)

	return upSQL, downSQL, nil
}

var postgresType = map[string]string{
	"key":     "SERIAL PRIMARY KEY",
	"int":     "INTEGER",
	"float64": "DOUBLE PRECISION",
	"text":    "TEXT",
	"string":  "VARCHAR(255)",
}

var sqlite3Types = map[string]string{
	"key":     "INTEGER PRIMARY KEY AUTOINCREMENT",
	"int":     "INTEGER",
	"float64": "REAL",
	"text":    "TEXT",
	"string":  "TEXT",
}

var mysqlTypes = map[string]string{
	"key":     "BIGINT AUTO_INCREMENT PRIMARY KEY",
	"int":     "INT",
	"float64": "DOUBLE",
	"text":    "TEXT",
	"string":  "VARCHAR(255)",
}

var migrationsTablePostgres = `
CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			migration_name VARCHAR(255) NOT NULL,
			applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`
var migrationsTableSqlite3 = `
CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			migration_name TEXT NOT NULL,
			applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

var migrationsTableMysql = `
CREATE TABLE IF NOT EXISTS migrations (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			migration_name VARCHAR(255) NOT NULL,
			applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`
