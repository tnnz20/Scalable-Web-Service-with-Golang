package biodata

import "fmt"

type Biodata struct {
	Name    string
	Address string
	Jobs    []string
	Reason  string
}

func (b Biodata) GetBiodata(number int) {
	fmt.Println("No Absen:", number)
	fmt.Println("============================================")
	fmt.Println("Nama:", b.Name)
	fmt.Println("Alamat:", b.Address)
	fmt.Println("Pekerjaan:", b.Jobs)
	fmt.Println("Alasan:", b.Reason)
}

func (b Biodata) GetListBiodata() []Biodata {
	return ListStudent
}
