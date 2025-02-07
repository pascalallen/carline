package main

import (
	"fmt"
	"github.com/pascalallen/carline/internal/carline/infrastructure/container"
	"github.com/pascalallen/carline/internal/carline/infrastructure/database/seeders"
)

func main() {
	c := container.InitializeContainer()
	defer c.MessageQueueConnection.Close()

	fmt.Println("Starting database seeding...")
	seeders.SeedDatabase(c.DatabaseSession)
	fmt.Println("Database seeding completed successfully.")
}
