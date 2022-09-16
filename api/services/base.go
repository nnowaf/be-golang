package services

import (
	"gobe/api/model"

	"github.com/jmoiron/sqlx"
)

type Service interface {
	PanggilKategoriX() []model.Kategori
	PanggilProdukX() []model.Produk
}

type service struct {
	db *sqlx.DB
}

func InitService(db *sqlx.DB) Service {
	return &service{db}
}
