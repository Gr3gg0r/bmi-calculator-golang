package main

import (
	"fmt"

	"github.com/gr3gg0r/bmi/info"
)

func main() {
	// Output information
	info.PrintWelcome()

	// Save that user input in variables
	weight := getWeight()
	height := getHeight()

	// Calculate the BMI  (weight / (height*height))
	bmi := calculateBMI(weight, height)

	fmt.Printf("Your BMI: %.2f", bmi)
}

func calculateBMI(weight float64, height float64) (bmi float64) {
	bmi = weight / (height * height)
	return
}
