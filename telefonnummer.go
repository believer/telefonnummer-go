package telefonnummer

import (
	"strings"
	"regexp"
)

func normalize(phoneNumber string) string {
	validNumbers := regexp.MustCompile(`\D`)

	if strings.Contains(phoneNumber, "(0)") {
		phoneNumber = strings.Replace(phoneNumber, "(0)", "", -1)
	}

	normalized := validNumbers.ReplaceAllString(phoneNumber, "")

	if normalized[:2] == "46" {
		return "0" + normalized[2:]
	}

	return normalized
}

func areaCodeDigitCount(phoneNumber string) int {
	doubleDigit := regexp.MustCompile(`^08`)
	tripleDigit := regexp.MustCompile(`^0(1[013689]|2[0136]|3[1356]|4[0246]|54|6[03]|7[0235-9]|9[09])`)

	if doubleDigit.MatchString(phoneNumber) {
		return 2
	} else if tripleDigit.MatchString(phoneNumber) {
		return 3
	}

	return 4
}

func makeRegex(areaCode int) func(int) *regexp.Regexp {
	return func(firstDigits int) *regexp.Regexp {
		if areaCode == 2 {
			if firstDigits == 3 {
				return regexp.MustCompile(`^(\d{2})(\d{3})(\d{2})(\d{2})$`)
			}

			if firstDigits == 10 {
				return regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{2})$`)
			}

			return regexp.MustCompile(`^(\d{2})(\d{2})(\d{2})(\d{2})$`)
		}

		if areaCode == 3 {
			if firstDigits == 3 {
				return regexp.MustCompile(`^(\d{3})(\d{3})(\d{2})(\d{2})$`)
			}

			return regexp.MustCompile(`^(\d{3})(\d{2})(\d{2})(\d{2})$`)
		}

		if areaCode == 4 {
			return regexp.MustCompile(`^(\d{4})(\d{2})(\d{2})(\d{2})$`)
		}

		return regexp.MustCompile(`^(\d{3})(\d{3})(\d{2})(\d{2})$`)
	}
}

// Parse - Parses a telephone number
func Parse(phoneNumber string) string {
	var result string

	normalized := normalize(phoneNumber)
	areaCode := areaCodeDigitCount(normalized)
	firstDigits := makeRegex(areaCode)

	voicemails := map[string]bool{
		"333": true,
		"222": true,
		"888": true,
	}

	if voicemails[normalized] {
		return "Röstbrevlåda"
	}

	switch areaCode {
	case 2:
		firstDigitsLength := 3

		if len(normalized) == 8 {
			firstDigitsLength = 2
		} else if len(normalized) == 10 {
			firstDigitsLength = 10
		}

		result = firstDigits(firstDigitsLength).ReplaceAllString(normalized, "$1-$2 $3 $4")
	case 3:
		firstDigitsLength := 3

		if len(normalized) == 9 {
			firstDigitsLength = 2
		}

		if len(normalized) == 8 {
			result = regexp.MustCompile(`^(\d{3})(\d{3})(\d{2})$`).ReplaceAllString(normalized, "$1-$2 $3")
		} else {
			result = firstDigits(firstDigitsLength).ReplaceAllString(normalized, "$1-$2 $3 $4")
		}
	case 4:
		if len(normalized) == 9 {
			result = regexp.MustCompile(`^(\d{4})(\d{3})(\d{2})$`).ReplaceAllString(normalized, "$1-$2 $3")
		} else {
			result = firstDigits(2).ReplaceAllString(normalized, "$1-$2 $3 $4")
		}
	}

	return result
}
