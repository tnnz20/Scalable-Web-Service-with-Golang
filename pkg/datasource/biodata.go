package datasource

import (
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/internal/domain"
)

var ListStudent = []domain.Biodata{
	{
		Name:    "Hakori",
		Address: "BJM",
		Jobs:    []string{"Student", "Programmer"},
		Reason:  "Suka Belajar",
	},
	{
		Name:    "Hanzo",
		Address: "JPG",
		Jobs:    []string{"Student"},
		Reason:  "Tidak tau yang lain",
	},
	{
		Name:    "Rakumen",
		Address: "ESP",
		Jobs:    []string{"Unkown"},
		Reason:  "Tidak tau menahu",
	},
	{
		Name:    "Leoti",
		Address: "BDG",
		Jobs:    []string{"Master"},
		Reason:  "Apalah arti sebuah alasan",
	},
}
