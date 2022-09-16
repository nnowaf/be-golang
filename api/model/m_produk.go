package model

type Produk struct {
	ProdukName string `json:"produkName"`
	ProdukKat  int    `json:"produkKat"`
	Kategori   `json:"katName"`
}
