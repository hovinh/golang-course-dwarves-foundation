package main

import (
	"fmt"
	"strconv"
)

func numDifferentIntegers(word string) int {
	str_digit_arr := parseStringDigit(word)
	// printStrArr(str_digit_arr)
	digit_decimal_arr := convertStringToDecimalDigit(str_digit_arr)
	// printIntArr(digit_decimal_arr)
	count := countUniqueInt(digit_decimal_arr)
	return count
}

func printStrArr(str_arr []string) {
	for _, ele := range(str_arr) {
		fmt.Printf("%s ", ele)
	}
	fmt.Println()
}

func printIntArr(int_arr []int) {
	for _, ele := range(int_arr) {
		fmt.Printf("%d ", ele)
	}
	fmt.Println()
}


func parseStringDigit(word string) []string {
	var str_digit_arr []string
	var temp_str string
	for _, char := range word {
		if ((char >= '0') && (char <= '9')){
			temp_str = temp_str + string([]rune{char})
		} else if (len(temp_str) > 0) {
			str_digit_arr = append(str_digit_arr, temp_str)
			temp_str = ""
		}
	}
	if (temp_str != "") {
		str_digit_arr = append(str_digit_arr, temp_str)
	}
	return str_digit_arr
}

func convertStringToDecimalDigit(str_digit_arr []string) []int {
	var digit_decimal_arr []int
	for _, str_digit := range str_digit_arr {
		// fmt.Println("Before", str_digit)
		str_digit_processed := removeLeadingZeros(str_digit)
		// fmt.Println("After", str_digit_processed)
		if (str_digit_processed == ""){
			continue
		}
		digit, _ := strconv.Atoi(str_digit_processed)
		digit_decimal_arr = append(digit_decimal_arr, digit)
	} 
	return digit_decimal_arr
}

func removeLeadingZeros(str string) string {
	prefix := 0
	for _, char := range str {
		if (char != '0') {
			break
		}
		prefix++
	}
	return str[prefix:]
}

func countUniqueInt(digit_decimal_arr []int) int {
	set := make(map[int]bool)
	for _, digit := range digit_decimal_arr {
		set[digit] = true
	}
	return len(set)
}

func main() {
	word := "a123bc34d8ef34"
	count := numDifferentIntegers(word)
	fmt.Printf("%v\n", count)

	word = "A1b01c001"
	count = numDifferentIntegers(word)
	fmt.Printf("%v\n", count)
}