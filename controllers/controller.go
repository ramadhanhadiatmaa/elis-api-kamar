package controllers

import (
	"kamar/models"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {

	var kamar []models.Kamar

	models.DB.Find(&kamar)

	return c.Status(fiber.StatusOK).JSON(kamar)

}

func Create(c *fiber.Ctx) error {

	var kamar models.Kamar
	if err := c.BodyParser(&kamar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&kamar).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Kamar berhasil dibuat",
	})
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")

	var kamar models.Kamar

	if err := c.BodyParser(&kamar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&kamar).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kamar tidak ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	var kamar models.Kamar

	if models.DB.Where("id = ?", id).Delete(&kamar).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Username tidak dapat dihapus.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus kamar.",
	})
}
