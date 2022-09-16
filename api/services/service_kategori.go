package services

import (
	"fmt"
	"gobe/api/model"
)

func (s service) PanggilKategoriX() []model.Kategori {
	var kategoriSlice []model.Kategori

	err := s.db.Select(
		&kategoriSlice,
		`SELECT category_id as KatId, 
		department_id as DeptId, 
		name as KatName, 
		description as Desc
		FROM public.category;`)

	if err != nil {
		fmt.Println(err.Error())
	}

	return kategoriSlice
}
