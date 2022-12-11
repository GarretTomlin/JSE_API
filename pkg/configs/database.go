package configs

import (
	"database/sql"
)

func ConnectToDB(driver string, dataSource string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
