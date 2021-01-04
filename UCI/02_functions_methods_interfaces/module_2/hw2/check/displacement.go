package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type InitialValues struct {
	acceleration float64
	initialVelocity float64
	intitalDisplacement float64
	time float64
}

func getNumericError() error {
	return errors.New("Please enter a numberic value")
}

func getUserInput(values *InitialValues) error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the acceleration:")
	scanner.Scan()
	userInput := scanner.Text()
	acceleration, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		err = getNumericError()
		return err
	}

	fmt.Println("Please enter the initial Velocity")
	scanner.Scan()
	userInput = scanner.Text()
	initalVelocity, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		err = getNumericError()
		return err
	}

	fmt.Println("Please enter the initial Displacement")
	scanner.Scan()
	userInput = scanner.Text()
	initialDisplacement, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		err = getNumericError()
		return err
	}

	fmt.Println("Please enter the time")
	scanner.Scan()
	userInput = scanner.Text()
	time, err := strconv.ParseFloat(userInput, 64)

	values.acceleration = acceleration
	values.initialVelocity = initalVelocity
	values.intitalDisplacement = initialDisplacement
	values.time = time
	return nil
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	fn:= func(time float64) float64 {
		return 0.5 * acceleration *math.Pow(time, 2.0) + initialVelocity * time + initialDisplacement
	}
	return fn
}

func main() {
	var initialValues InitialValues
	err := getUserInput(&initialValues)
	if err != nil {
		panic(getNumericError())
	}
	DisplaceFn := GenDisplaceFn(initialValues.acceleration, initialValues.initialVelocity,
		initialValues.intitalDisplacement)
	fmt.Printf("Displacement at time %f = %f", initialValues.time, DisplaceFn(initialValues.time))
}