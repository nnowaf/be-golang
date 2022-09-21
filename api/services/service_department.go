package services

import (
	"fmt"
	"gobe/api/model"
)

func (s service) PostDepartmentX(data model.Department) {

	_, err := s.db.Query(`INSERT INTO public.department(
		name, description)
		VALUES ($1, $2);`, data.Name, data.Descdep)
	if err != nil {
		fmt.Println(err.Error())
	}
}
