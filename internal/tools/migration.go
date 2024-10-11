package tools

import "path/filepath"

// MigrationDir provides relative path to migration dir
func PathMigrationDir() string {
	return filepath.Join("internal", "database", "migrations")
}
