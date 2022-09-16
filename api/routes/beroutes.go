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
	app.Get("/api/kategori/list", ui.ListKategori(s))
	app.Get("/api/produk/list", ui.ListProduk(s))
}
