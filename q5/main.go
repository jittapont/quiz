package main

import (
	"fmt"
	"math"
)

func calculateParkingFee(hour, minute int) int {
	minutes := hour*60 + minute
	cost := 0
	if minutes > 30 && minutes <= 180 {
		minutes -= 30
		div := float64(minutes) / 60
		if minutes/60 > 0 {
			cost += (minutes / 60) * 10
		} else {
			cost += 10
		}
		if div > 1 {
			cost += 20
		}

	}
	if minutes > 180 && minutes <= 360 {
		cost += 30
		minutes -= 210
		div := int(math.Ceil(float64(minutes) / 60))
		if div > 0 {
			cost += div * 20
		} else {
			cost += 20
		}

	}
	if minutes > 360 {
		cost += 90
		minutes -= 360
		div := minutes / 1440
		if div > 0 {
			cost += div * 200
		}
	}
	return cost
}

func main() {
	fmt.Println(calculateParkingFee(4, 15))
	fmt.Println(calculateParkingFee(2, 10))
}
