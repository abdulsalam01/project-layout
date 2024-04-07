package tools

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func SchemaMigrate(connUrl string, version uint) error {
	var (
		err error
	)

	migrations, err := migrate.New("file://config/database/migrations", connUrl)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil {
		return err
	}

	defer migrations.Close()
	return nil
}
