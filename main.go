package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"crud-api/controllers" 
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func main() {
	
	database, err := gorm.Open(sqlite.Open("data.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed ...........")
	}
	db = database

	
	db.AutoMigrate(&Product{})

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db) 
		return c.Next()
	})
	

	
	app.Get("/products", controllers.GetAllProducts) 
	app.Get("/products/:id", controllers.GetProduct)
	app.Post("/products", controllers.CreateProduct)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)


	app.Listen(":3000")
}
