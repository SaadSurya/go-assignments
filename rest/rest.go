package rest

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/saadsurya/go-assignment/rest/book"
	"github.com/saadsurya/go-assignment/rest/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=go_assignment sslmode=disable password=10pearls1+")
	if err != nil {
		panic("Failed to connect to database")
	} 
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func Main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
