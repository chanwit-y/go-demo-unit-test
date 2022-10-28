//go:generate go run db_migrate.go
package main

import "go-demo-unit-test/pkg/db"

func main() {
	db.Migrate()
}
