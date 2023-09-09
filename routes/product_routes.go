package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"crud-api/controllers"
)


func DefineRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/products", controllers.GetAllProducts)
	app.Get("/products/:id", controllers.GetProduct)
	app.Post("/products", controllers.CreateProduct)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)
}
