package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Welcome to Chess Bishop!")
	// Read 2 strings from the user
	var input1, input2 string
	validColumns := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	validRows := []string{"1", "2", "3", "4", "5", "6", "7", "8"}

	for {
		fmt.Print("Enter the first bishop position: ")
		fmt.Scanln(&input1)
		fmt.Print("Enter the second bishop position: ")
		fmt.Scanln(&input2)

		input1 = strings.ToUpper(input1)
		input2 = strings.ToUpper(input2)

		fmt.Println("You entered: ", input1, input2)

		if validInput(input1, validColumns, validRows) &&
			validInput(input2, validColumns, validRows) {
			break
		}
		fmt.Println("Invalid input. Please enter a valid chess position (e.g. A1)")
	}
	fmt.Printf("You entered: %s and %s\n", input1, input2)

	canAttackEachOther := canAttackEachOther(input1, input2)
	fmt.Printf("The bishops can attack each other: %t\n", canAttackEachOther)
}

func validInput(input string, validColumns []string, validRows []string) bool {
	if len(input) == 2 {
		colInput := string(input[0])
		rowInput := string(input[1])
		if slices.Contains(validColumns, colInput) &&
			slices.Contains(validRows, rowInput) {
			return true
		}
	}
	return false
}

func canAttackEachOther(input1, input2 string) bool {
	onSameRow := onSameRow(input1, input2)
	onSameColumn := onSameColumn(input1, input2)
	onSameDiagonal := onSameDiagonal(input1, input2)
	if onSameDiagonal {
		return true
	}
	//kinda redundant
	if onSameRow || onSameColumn {
		return false
	}
	return false
}

func onSameDiagonal(input1, input2 string) bool {
	//chess notation so A1 and B2 are on the same diagonal
	col1 := int(input1[0])
	row1 := int(input1[1])
	col2 := int(input2[0])
	row2 := int(input2[1])
	deltaCols := math.Abs(float64(col1 - col2))
	deltaRows := math.Abs(float64(row1 - row2))
	return deltaCols == deltaRows
}

func onSameColumn(input1, input2 string) bool {
	//chess notation so A1 and A2 are on the same column
	return input1[0] == input2[0]
}

func onSameRow(input1, input2 string) bool {
	//chess notation so A1 and B1 are on the same row
	return input1[1] == input2[1]
}
