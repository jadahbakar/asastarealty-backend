package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/jadahbakar/asastarealty-backend/internal/infrastructure/postgresql"
	"github.com/jadahbakar/asastarealty-backend/internal/infrastructure/version"
	"github.com/jadahbakar/asastarealty-backend/pkg/utils"
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
	log.Printf("Defining Config.....................")
	config, err := config.New(".")
	if err != nil {
		log.Printf("error Loading Config -> %v\n", err)
	}
	// Define Database
	log.Printf("Defining Database...................")
	dbClient, err := postgresql.NewPgClient(config)
	if err != nil {
		log.Panicf("Error Connecting Database -> %v\n", err)
	}
	// Define Server
	log.Printf("Defining Apps.......................")
	apps := app.New(config, dbClient)
	engine := apps.GetEngine()
	logger := apps.GetLogger()

	// Start the Server
	utils.StartFiberWithGracefulShutdown(engine, dbClient, config, logger)
}
