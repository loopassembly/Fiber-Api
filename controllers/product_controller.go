package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"crud-api/models"
)


func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	db := c.Locals("db").(*gorm.DB) 
	db.Find(&products)
	return c.JSON(products)
}


func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	db := c.Locals("db").(*gorm.DB) 
	db.First(&product, id)
	return c.JSON(product)
}


func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	db := c.Locals("db").(*gorm.DB) 
	db.Create(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	db := c.Locals("db").(*gorm.DB) 
	if err := db.First(&product, id).Error; err != nil {
		return err
	}
	updatedProduct := new(models.Product)
	if err := c.BodyParser(updatedProduct); err != nil {
		return err
	}
	db.Model(&product).Updates(updatedProduct)
	return c.JSON(product)
}


func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	db := c.Locals("db").(*gorm.DB) 
	if err := db.First(&product, id).Error; err != nil {
		return err
	}
	db.Delete(&product)
	return c.SendStatus(fiber.StatusNoContent)
}
