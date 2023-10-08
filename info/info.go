package info

import (
	"bufio"
	"fmt"
)

const mainTitle = "BMI Calculator"
const seperator = "------------------------"
const weightPrompt = "Please enter your weight (kg): "
const heightPrompt = "Please enter your height (m): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(seperator)
}

func GetWeight(reader bufio.Reader) string {
	fmt.Print(weightPrompt)
	weightInput, _ := reader.ReadString('\n')
	return weightInput
}

func GetHeight(reader bufio.Reader) string {
	fmt.Print(heightPrompt)
	heighInput, _ := reader.ReadString('\n')
	return heighInput
}
