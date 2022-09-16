package model

type Kategori struct {
	KatId   int    `json:"katId"`
	DeptId  int    `json:"deptId"`
	KatName string `json:"katName"`
	Desc    string `json:"desc"`
}
