package data

import (
	"fmt"
	"strings"
)

func getAll(table string) string {
	return fmt.Sprintf("SELECT * FROM %s", table)
}

func getById(table string) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table)
}

func insert(table string, columns []string) string {
	values := []string{}
	for i := 1; i <= len(columns); i++ {
		values = append(values, fmt.Sprintf("$%d", i))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, strings.Join(columns, ", "), strings.Join(values, ", "))
}

// func delete() string {

// }

// func update() string {

// }
