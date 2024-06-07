package main

import (
	"fmt"
	"math"
	"os"
)

var (
	studentRange [2]int = [2]int{1, 60}
	gradeRange   [2]int = [2]int{0, 100}
)

const (
	exitCommand string = "e"
)

func main() {
	num := setTotalStudents()

	grades := setGrades(num)

	getFinalGrade(grades)
}

func setInput() (num int) {
	var input string

	fmt.Scan(&input)

	if input == "e" {
		os.Exit(0)
	}

	fmt.Sscanf(input, "%d", &num)

	return num
}

func setTotalStudents() (num int) {

	fmt.Printf("\nInput total students : ")

	num = setInput()

	if num < studentRange[0] || num > studentRange[1] {
		fmt.Printf("\nInvalid range of input, valid range %v - %v, or '%v' to exit", studentRange[0], studentRange[1], exitCommand)

		return setTotalStudents()
	}

	return num
}

func setGrades(numberOfStudents int) []int {
	grades := make([]int, numberOfStudents)

	for i := 0; i < numberOfStudents; {
		fmt.Printf("\nInput Grade for student number %v: ", i+1)

		num := setInput()

		if num < gradeRange[0] || num > gradeRange[1] {
			fmt.Printf("\nInvalid range of input, valid range %v - %v, or '%v' to exit", gradeRange[0], gradeRange[1], exitCommand)
		} else {
			grades[i] = num

			i++
		}
	}

	return grades
}

func getFinalGrade(grades []int) {
	for i, g := range grades {

		if g < 40 {
			fmt.Printf("\nStudent %v: Fail", i+1)
			continue
		}

		if g%5 > 2 {
			fmt.Printf("\nFinal Grade for student %v : %v", i+1, math.Ceil(float64(g)/5)*5)
		} else {
			fmt.Printf("\nFinal Grade for student %v : %v", i+1, g)
		}
	}
}
