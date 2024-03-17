package helpers

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/models"
)

func updateElement(min, max int) (result *models.Element) {
	result = &models.Element{
		Water: rune(rand.Intn(max-min) + min),
		Wind:  rune(rand.Intn(max-min) + min),
	}
	return
}

func checkCondition(element *models.Element) (status *models.Status) {
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

	status = &models.Status{
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}
	return
}

func RunCronJob() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(15).Seconds().Do(func() {
		Element := updateElement(1, 100)
		Condition := checkCondition(Element)

		log.Printf("Element Water: %v m, Status %v\n", Element.Water, Condition.WaterStatus)
		log.Printf("Element Wind: %v m/s, Status %v\n", Element.Wind, Condition.WindStatus)
	})

	s.StartBlocking()
}
