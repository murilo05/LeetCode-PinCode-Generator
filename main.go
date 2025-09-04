package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	//Quantity of PinCodes to Generate
	input := 100

	pinCodeList := generatePinCodeList(input)

	fmt.Println(pinCodeList)
}

// generatePinCodeList - Generate a list of pincodes, the quantity can be insert in the param
func generatePinCodeList(quantity int) map[int]string {
	listOfPinCodes := make(map[int]string)
	seen := make(map[string]bool)

	for j := 0; j <= quantity; j++ {
		pinCode := generatePinCode()
		key := convertPinToString(pinCode)

		//Check if exists a PinCode with existing numbers
		if !seen[key] {
			listOfPinCodes[j] = key
			seen[key] = true
		}
	}
	return listOfPinCodes
}

// generatePinCode - Generate a 4 digits pinCode
func generatePinCode() []int {
	var slice []int
	for j := 0; j < 4; j++ {
		slice = append(slice, rand.Intn(9))
	}

	return slice
}

// convertPinToString - Convert the PinCode to string
func convertPinToString(pinCode []int) string {
	var pinCodeStr string
	for _, d := range pinCode {
		pinCodeStr += strconv.Itoa(d)
	}
	return pinCodeStr
}
