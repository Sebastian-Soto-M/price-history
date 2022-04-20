package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name   string
	Domain string
}

func (store *Store) Details() string {
	return fmt.Sprintf("%-6s | %15s | Tracked Items = 0", store.Name, store.Domain)
}
