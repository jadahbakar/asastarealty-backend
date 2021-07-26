package postgresql

import (
	"database/sql"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
)

func NewPgClient(config *config.Config) (*sql.DB, error) {
	connConfig, _ := pgx.ParseConfig(config.Db.Url)
	connStr := stdlib.RegisterConnConfig(connConfig)
	dbConn, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error Loading DB Connection -> %v\n", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Error Loading DB Ping -> %v\n", err)
	}
	return dbConn, err
}
