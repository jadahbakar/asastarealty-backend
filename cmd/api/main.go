package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	// new

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/pkg/config"
	"github.com/jadahbakar/asastarealty-backend/pkg/utils"
	"github.com/jadahbakar/asastarealty-backend/pkg/version"
)

// https://gist.github.com/rnyrnyrny/282fe705d6e8dc012e482582d7c8ec0b
func init() {
	ver := flag.Bool("version", false, "print version information")
	v := flag.Bool("v", false, "print version information")

	flag.Parse()
	if *v || *ver {
		fmt.Println(version.String())
		os.Exit(0)
	}
}

func main() {
	// Define Config
	log.Printf("Defining Config....")
	config, err := config.New()
	if err != nil {
		log.Printf("error Loading Config -> %v\n", err)
	}

	log.Printf("Defining Database....")

	connConfig, _ := pgx.ParseConfig(config.DbUrl)
	connStr := stdlib.RegisterConnConfig(connConfig)
	dbConn, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Define Server
	log.Printf("Defining Apps....")
	apps := app.New(config, dbConn)
	engine := apps.GetEngine()
	logger := apps.GetLogger()

	// Start the Server
	utils.StartFiberWithGracefulShutdown(engine, dbConn, config, logger)
}
