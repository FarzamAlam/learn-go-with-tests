package property

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumeral = []RomanNumeral{
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

func ConvertToRoman(arabic int) string {
	var res strings.Builder
	for _, numeral := range allRomanNumeral {
		for arabic >= numeral.Value {
			res.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return res.String()
}

func ConvertToArabic(roman string) int {
	return 1
}
