package main

import (
	"fmt"
	"gobe/api/routes"
	"gobe/api/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "123456", "bootcamp")

	dbx, err := sqlx.Connect("postgres", connection)
	if err != nil {
		log.Fatalln(err)
	}
	s := services.InitService(dbx)
	routes.BeRouter(app, s)
	app.Listen(":26")
}
