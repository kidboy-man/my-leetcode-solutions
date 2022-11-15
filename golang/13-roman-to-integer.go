package main

func romanToInt(s string) int {
	mapping := make(map[string]int)
	mapping["I"] = 1
	mapping["V"] = 5
	mapping["X"] = 10
	mapping["L"] = 50
	mapping["C"] = 100
	mapping["D"] = 500
	mapping["M"] = 1000

	var total int
	var charBefore string
	for _, char := range s {
		strChar := string(char)
		total += mapping[strChar]
		switch charBefore {
		case "I":
			if strChar == "V" || strChar == "X" {
				total -= 2
			}
		case "X":
			if strChar == "L" || strChar == "C" {
				total -= 20
			}
		case "C":
			if strChar == "D" || strChar == "M" {
				total -= 200
			}
		}
		charBefore = strChar
	}

	return total
}
