package tests

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Using a pre defined test database
// Make sure you have ran make migratetestup
func GetTestDB() *gorm.DB {
	dbUrl := "postgres://matawis_user:matawis_pass@localhost:5432/matawis_test"
	TestDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to load config")
	}

	return TestDB
}

func GetExistingTestDBTables(sqlDB *sql.DB) []string {
	tables := make([]string, 0)

	rows, err := sqlDB.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Loop over the rows and append the table names to the tables array
	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			panic(err.Error())
		}
		tables = append(tables, tableName)
	}
	return tables
}

func TruncateDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	tableNames := GetExistingTestDBTables(sqlDB)
	// Truncate each table
	for _, tableName := range tableNames {
		if err := db.Exec("TRUNCATE TABLE " + tableName + " CASCADE;").Error; err != nil {
			return err
		}
	}

	return nil
}

func GetMockDB(t *testing.T) (gdb *gorm.DB, mock sqlmock.Sqlmock) {
	// Create a mock MySQL database and Gorm instance
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}
	gdb, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to create Gorm instance: %s", err)
	}

	return gdb, mock
}
