package solve

import (
	"math"
	"strings"
)

func TotalAmount(productName string, amount float64, phoneNumber string, duration int) float64 {
	const (
		PHONE      = "phone"
		COMPUTER   = "computer"
		TELEVISION = "television"
	)

	productName = strings.TrimSpace(strings.ToLower(productName))

	if amount < 0 || duration < 0 || productName == "" {
		return 0
	}

	if productName == PHONE {
		return totalAmountPhone(amount, duration)
	} else if productName == COMPUTER {
		return totalAmountComputer(amount, duration)
	} else if productName == TELEVISION {
		return totalAmountTelevision(amount, duration)
	} else {
		return 0
	}
}

func totalAmountPhone(amount float64, duration int) float64 {
	rangeMax := 9
	percent := 3
	unitCount := int(math.Ceil(float64(duration) / float64(rangeMax)))
	total := float64(unitCount*percent) / 100.0
	return amount + amount*total
}

func totalAmountComputer(amount float64, duration int) float64 {
	rangeMax := 12
	percent := 4
	unitCount := int(math.Ceil(float64(duration) / float64(rangeMax)))
	total := float64(unitCount*percent) / 100.0
	return amount + amount*total
}

func totalAmountTelevision(amount float64, duration int) float64 {
	rangeMax := 18
	percent := 5
	unitCount := int(math.Ceil(float64(duration) / float64(rangeMax)))
	total := float64(unitCount*percent) / 100.0
	return amount + amount*total
}
