package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalf("invalid port number: %s", err.Error())
	}

	fmt.Printf("Server started in port %d\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
