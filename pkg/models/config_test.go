package models

import (
	"os"
	"testing"

	"github.com/matawis/matawis/tests"
	"gorm.io/gorm"
)

var (
	TestDB *gorm.DB
)

func TestMain(m *testing.M) {
	TestDB = tests.GetTestDB()
	tests.TruncateDatabase(TestDB)
	os.Exit(m.Run())

}
