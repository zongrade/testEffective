// Package tools provides utils
package tools

import (
	"errors"
	"os"
)

// ErrorMigrationDirNotExist says have problem with path to migration dir.
var ErrMigrationDirNotExist = errors.New("директори миграции не существует или путь не верен")

// IsMigrationDirExist check migration dir is exist.
func IsMigrationDirExist() error {
	migrationsDir := PathMigrationDir()

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return ErrMigrationDirNotExist
	}

	return nil
}
