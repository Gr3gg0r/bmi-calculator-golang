package main

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gr3gg0r/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func sanitizeValueByPlatform(input string) string {
	if runtime.GOOS == "windows" {
		return strings.Replace(input, "\r\n", "", -1)
	}
	return strings.Replace(input, "\n", "", -1)
}

func convertStringToFloat(input string) (i float64) {
	result := sanitizeValueByPlatform(input)
	i, _ = strconv.ParseFloat(result, 64)
	return
}

func getWeight() (r float64) {
	weightInput := info.GetWeight(*reader)
	r = convertStringToFloat(weightInput)
	return
}

func getHeight() (r float64) {
	heightInput := info.GetHeight(*reader)
	r = convertStringToFloat(heightInput)
	return
}
