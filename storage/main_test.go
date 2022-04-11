package storage

import (
	"fmt"
	"os"
	"testing"

	"github.com/uptrace/bun"

	testutilsPostgres "ikea/testutils/postgres"
)

var db *bun.DB

func InitTestDB() (*bun.DB, func()) {
	db, dbCleaner := testutilsPostgres.Setup("../scripts/migrations")
	db.AddQueryHook(CustomQueryHook())
	return db, dbCleaner
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) (res int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			res = 1
		}
	}()

	var cleaner func()
	db, cleaner = InitTestDB()
	defer cleaner()

	res = m.Run()
	return res
}
