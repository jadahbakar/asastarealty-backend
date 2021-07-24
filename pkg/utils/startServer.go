package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/pkg/config"
)

func StartFiberWithGracefulShutdown(fiberApp *fiber.App, db *sql.DB, config *config.Config, logFile *os.File) {
	// Listen from a different goroutine
	go func() {
		if err := fiberApp.Listen(fmt.Sprintf(":%d", config.AppPort)); err != nil {
			messageErr := fmt.Sprintf("Server is not running! on Port %d Reason: %v", config.AppPort, err)
			log.Panic(messageErr)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	_ = <-c                                         // This blocks the main thread until an interrupt is received
	// fmt.Println("Gracefully shutting down...")
	log.Printf("Gracefully shutting down...")
	_ = fiberApp.Shutdown()
	// fmt.Println("Running cleanup tasks...")
	log.Printf("Running cleanup tasks...")
	// Your cleanup tasks go here
	db.Close()
	logFile.Close()
	// redisConn.Close()
	// fmt.Println("Fiber was successful shutdown.")
	log.Println("Fiber was successful shutdown.")

}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App, config *config.Config, logFile *os.File) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
		log.Println("Shutting down server...")

	}()

	// Run server.
	if err := a.Listen(fmt.Sprintf(":%d", config.AppPort)); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
	// Clossing the Fiber Log File
	defer logFile.Close()
	log.Println("Close log file...")

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App, config *config.Config) {
	// Run server.
	if err := a.Listen(fmt.Sprintf(":%d", config.AppPort)); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
