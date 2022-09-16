package services

import (
	"fmt"
	"gobe/api/model"
)

func (s service) PanggilProdukX() []model.Produk {
	var produkSlice []model.Produk

	err := s.db.Select(
		&produkSlice,
		`SELECT product.name as ProdukName,
    	product_category.category_id as ProdukKat,
    	category.name as KatName
		FROM product
		LEFT JOIN product_category ON product_category.product_id = product.product_id
		LEFT JOIN category ON category.category_id = product_category.category_id;`)

	if err != nil {
		fmt.Println(err.Error())
	}

	return produkSlice
}
