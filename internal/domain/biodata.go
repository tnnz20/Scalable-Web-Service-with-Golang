package domain

import "fmt"

type Biodata struct {
	Name    string
	Address string
	Jobs    []string
	Reason  string
}

func (b Biodata) PrintBiodata(number int) {
	fmt.Printf("\nNo Absen: %v", number)
	fmt.Println("============================================")
	fmt.Println("Nama: ", b.Name)
	fmt.Println("Alamat: ", b.Address)
	fmt.Println("Pekerjaan: ", b.Jobs)
	fmt.Println("Alasan Memilih Golang: ", b.Reason)
}
