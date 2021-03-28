package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var errNIsNegative = errors.New("n can't be negative number")

func main() {
	sum, err := calculateSumUntil(3)
	if err != nil {
		log.Fatal("error in calculate sum: ", err)
	}
	fmt.Println(sum)
}

func calculateSumUntil(n int) (int, error) {
	sum := 0
	if n < 0 {
		return 0, errNIsNegative
	}
	for {
		if n == 0 {
			break
		}
		sum += int(math.Pow(float64(n), float64(n)))
		n -= 1
	}
	return sum, nil
}
