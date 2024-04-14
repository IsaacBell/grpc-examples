package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseAdapter struct {
	db *gorm.DB
}

func NewDatabaseAdapter(conn *sql.DB) (*DatabaseAdapter, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Can't connect database (gorm) : %v", err)
	}

	return &DatabaseAdapter{
		db: db,
	}, nil
}
