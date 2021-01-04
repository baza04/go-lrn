package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	requests := []string{"acceleration", "velocity", "start-point", "time"}
	p := map[string]float64{}

	for _, parameter := range requests {
		p[parameter] = ReadParameters(parameter)
	}
	calcDiplacement := GenDisplaceFn(p["acceleration"], p["velocity"], p["start-point"])

	fmt.Printf("diplacement after %.1f second will equal: %.2f\n", p["time"], calcDiplacement(p["time"]))
}

// ReadParameters return input after converting to float
func ReadParameters(parameter string) float64 {
	fmt.Printf("Enter %s: ", parameter)
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	floatInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("!!! ERROR !!! Incorrect input of %s: \"%v\"\nInput value must be float64 type\n", parameter, input)
		os.Exit(1)
	}
	return floatInput
}

// GenDisplaceFn return func to calculate deplacement
func GenDisplaceFn(acceleration, velocity, start float64) func(float64) float64 {
	Fn := func(time float64) float64 {
		return start + acceleration*math.Pow(time, 2)/2 + velocity*time
	}
	return Fn
}
