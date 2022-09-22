package services

import (
	"fmt"
	"gobe/api/model"
)

func (s service) PostDepartmentX(data model.Department) (model.Department, error) {

	_, err := s.db.Query(`INSERT INTO public.department(
		name, description)
		VALUES ($1, $2);`, data.Name, data.Descdep)
	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}

	return data, err
}

func (s service) PanggilDeptXById(id int) model.Department {
	var deptSlice model.Department

	err := s.db.Get(&deptSlice, `SELECT department_id as Deptid, coalesce("name", '') as Name , coalesce(description, '-') as Descdep FROM department where department_id = $1`, id)

	if err != nil {
		fmt.Println(err.Error())
	}

	return deptSlice
}
