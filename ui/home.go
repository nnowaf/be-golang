package ui

import (
	"gobe/api/services"

	"github.com/gofiber/fiber/v2"
)

func Home() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("KODOK HOME")
	}
}

func About() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("KODOK ABOUT")
	}
}

func Nilai() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, " + c.Params("nilai"))
	}
}

func ListKategori(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := s.PanggilKategoriX()
		return c.JSON(data)
	}
}

func ListProduk(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := s.PanggilProdukX()
		return c.JSON(data)
	}
}
