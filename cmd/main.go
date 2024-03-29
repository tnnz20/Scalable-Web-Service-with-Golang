package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/pkg/datasource"
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/pkg/helpers"
)

func main() {
	Students := datasource.ListStudent

	if len(os.Args) > 1 {
		input, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Bukan angka, mohon masukan angka...")
		}

		// Check Input
		bio, err := helpers.CheckInput(input, Students)
		if err != nil {
			panic(err)
		}

		bio.PrintBiodata(input)
	} else {
		panic("Mohon masukan angka terlebih dahulu...")
	}
	fmt.Println("\nProgram Selesai...")
}
