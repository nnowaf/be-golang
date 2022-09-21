package routes

import (
	"gobe/api/services"
	"gobe/ui"

	"github.com/gofiber/fiber/v2"
)

func BeRouter(app fiber.Router, s services.Service) {
	app.Get("/", ui.Home())
	app.Get("/about", ui.About())
	app.Get("/nilai/:nilai", ui.Nilai())

	app.Get("/api/produk/list", ui.ListProduk(s))

	app.Get("/api/kategori/list", ui.ListKategori(s))
	app.Get("/api/department/list", ui.ListDepartment(s))

	app.Post("/api/kategori/post", ui.ListKategoriTable(s))
	app.Post("/api/department/post", ui.PostDepartment(s))

}
