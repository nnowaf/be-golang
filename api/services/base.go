package services

import (
	"gobe/api/model"

	"github.com/jmoiron/sqlx"
)

type Service interface {
	PanggilKategoriX(page model.Halaman) []model.Kategori
	PanggilDepartmentX() []model.Department
	PanggilProdukX() []model.Produk
	BanyakKategoriX(page model.Halaman) int64
	PostDepartmentX(data model.Department)
}

type service struct {
	db *sqlx.DB
}

func InitService(db *sqlx.DB) Service {
	return &service{db}
}
