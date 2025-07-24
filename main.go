package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	port := ":9000"
	e := echo.New()
	
	if err := e.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
