package ui

import (
	"gobe/api/model"
	"gobe/api/services"
	"gobe/api/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
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

func ListProduk(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := s.PanggilProdukX()
		return c.JSON(data)
	}
}

func ListKategori(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.Halaman
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		data := s.PanggilKategoriX(req)
		return c.JSON(data)
	}
}

func ListKategoriTable(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.Halaman
		var output model.Pagedata

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		output.Totalelements = s.BanyakKategoriX(req)
		output.Pagenumber = req.Pagenumber
		output.Size = req.Size
		output.Data = s.PanggilKategoriX(req)
		return c.JSON(output)
	}
}

func PostDepartment(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req model.Department

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		var validate = validator.New()
		err := validate.Struct(req)
		if err != nil {
			return c.Status(405).JSON(utils.ValidateStruct(err))
		}
		data, err := s.PostDepartmentX(req)
		if err != nil {
			return c.Status(405).JSON(err)
		}

		return c.JSON(data)
	}
}

func ListDepartment(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data := s.PanggilDepartmentX()
		return c.JSON(data)
	}
}

func GetDeptById(s services.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idDept, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		data := s.PanggilDeptXById(idDept)
		return c.JSON(data)
	}
}
