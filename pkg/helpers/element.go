package helpers

import (
	"math/rand"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/internal/domain"
)

func UpdateElement(min, max int) (result *domain.Element) {
	result = &domain.Element{
		Water: rune(rand.Intn(max-min) + min),
		Wind:  rune(rand.Intn(max-min) + min),
	}
	return
}

func CheckElement(element *domain.Element) (status *domain.Status) {
	var waterStatus string
	var windStatus string

	// Water Condition
	switch {
	case element.Water <= 5:
		waterStatus = "Aman"
	case element.Water >= 6 && element.Water <= 8:
		waterStatus = "Siaga"
	case element.Water > 8:
		waterStatus = "Bahaya"
	default:
		waterStatus = "Water Value Not Defined"
	}

	// Wind Condition
	switch {
	case element.Wind <= 6:
		windStatus = "Aman"
	case element.Wind >= 7 && element.Wind <= 15:
		windStatus = "Siaga"
	case element.Wind > 15:
		windStatus = "Bahaya"
	default:
		windStatus = "Wind Value Not Defined"
	}

	status = &domain.Status{
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}
	return
}
