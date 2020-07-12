package property

import "strings"

func ConvertToRoman(arabic int) string {
	var res strings.Builder
	for arabic > 0 {
		switch {
		case arabic > 9:
			res.WriteString("X")
			arabic -= 10
		case arabic > 8:
			res.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			res.WriteString("V")
			arabic -= 5
		case arabic > 3:
			res.WriteString("IV")
			arabic -= 4
		default:
			res.WriteString("I")
			arabic--
		}
	}
	return res.String()
}
