package helpers

import (
	"errors"
	"fmt"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/internal/domain"
)

func CheckInput(input int, biodatas []domain.Biodata) (*domain.Biodata, error) {
	if input > len(biodatas) || input < 1 {
		msg := fmt.Sprintf("INDEX OUT OF RANGE: Mohon Masukan Angka dari 1 - %d", len(biodatas))
		return nil, errors.New(msg)
	} else {
		return &biodatas[input-1], nil
	}
}
