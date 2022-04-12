package main

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToArabic(roman string) int {
	total := 0
	for range roman {
		total++
	}
	return total
}

func ConvertToRoman(number int) string {
	var result strings.Builder //builder jest uzywany do budowania stringow przy uzyciu metody write

	for _, numeral := range allRomanNumerals {
		for number >= numeral.Value {
			result.WriteString(numeral.Symbol)
			number -= numeral.Value
		}
	}

	/* wersja 2
	for number > 0 {
		switch {
		case number > 9:
			result.WriteString("X")
			number -= 10
		case number > 8:
			result.WriteString("IX")
			number -= 9
		case number > 4:
			result.WriteString("V")
			number -= 5
		case number > 3:
			result.WriteString("IV")
			number -= 4
		default:
			result.WriteString("I")
			number--
		}
	}
	*/
	/* wersja 1
	for i := 0; i < number; i++ {
		if number == 5 {
			result.WriteString("V")
			break
		}
		if number == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}
	*/
	return result.String()
}
