package services

import (
	"fmt"
	"gobe/api/model"
	"gobe/api/utils"
)

func (s service) PanggilDepartmentX() []model.Department {
	var departMenSlice []model.Department

	err := s.db.Select(&departMenSlice, `SELECT department_id as Deptid, name as Name, coalesce(description, '') as Descdep from department`)
	if err != nil {
		fmt.Println(err.Error())
	}
	return departMenSlice
}

func (s service) BanyakKategoriX(page model.Halaman) (data int64) {
	var where string
	if page.Param != nil && page.Param["nama"] != "" && page.Param["dpdown"] != "" {
		where = " and name ilike '%' ||'" + page.Param["nama"] + "'||'%' and department_id = '" + page.Param["dpdown"] + "'"
	}

	err := s.db.Get(&data, `SELECT count(category_id) FROM category WHERE 1=1`+where+`    `)

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return data
}

func (s service) PanggilKategoriX(page model.Halaman) []model.Kategori {
	var kategoriSlice []model.Kategori

	var where string

	if page.Param != nil && page.Param["nama"] != "" && page.Param["dpdown"] != "" {
		where = " and name ilike '%' ||'" + page.Param["nama"] + "'||'%' and department_id = '" + page.Param["dpdown"] + "'"
	}

	err := s.db.Select(
		&kategoriSlice,
		`SELECT category_id as KatId, 
		department_id as DeptId, 
		name as KatName, 
		description as Desc
		FROM public.category where 1=1 `+where+` limit $1 offset $2;`, page.Size, utils.GetEnd(page.Pagenumber, page.Size))

	if err != nil {
		fmt.Println(err.Error())
	}

	return kategoriSlice
}
