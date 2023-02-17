package restapi_test

import (
	"database/sql"

	"github.com/grailbio-external/goose"
)

func init() {
	goose.AddMigration(Up, Down)
}

func Up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS info2;")

	if err != nil {
		return err
	}
	return nil
}

func Down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE info2;")
	if err != nil {
		return err
	}
	return nil
}
