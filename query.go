package data

import (
	"fmt"
)

func getAll(table string) string {
	return fmt.Sprintf("SELECT * FROM %s", table)
}

func getById(table string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table)
}
