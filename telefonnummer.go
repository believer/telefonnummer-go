package telefonnummer

import (
	"regexp"
	"strconv"
	"strings"
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

func convertToString(number int) string {
	return strconv.Itoa(number)
}

func makeRegex(areaCode int) func(int) *regexp.Regexp {
	return func(firstDigits int) *regexp.Regexp {
		ac := convertToString(areaCode)
		digits := convertToString(firstDigits)
		standardFormat := "^(\\d{" + ac + "})(\\d{" + digits + "})(\\d{2})(\\d{2})$"

		if firstDigits == 10 {
			return regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{2})$`)
		}

		return regexp.MustCompile(standardFormat)
	}
}

func shortFormat(areaCode int, phoneNumber string) string {
	ac := convertToString(areaCode)
	shortRegex := "^(\\d{" + ac + "})(\\d{3})(\\d{2})$"

	return regexp.MustCompile(shortRegex).ReplaceAllString(phoneNumber, "$1-$2 $3")
}

func normalFormat(regex *regexp.Regexp, phoneNumber string) string {
	return regex.ReplaceAllString(phoneNumber, "$1-$2 $3 $4")
}

// Parse - Parses a telephone number
func Parse(phoneNumber string) string {
	normalized := normalize(phoneNumber)
	areaCode := areaCodeDigitCount(normalized)
	firstDigits := makeRegex(areaCode)
	numberLength := len(normalized)
	firstDigitsLength := 3

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
		if numberLength == 8 {
			firstDigitsLength = 2
		} else if numberLength == 10 {
			firstDigitsLength = 10
		}

		return normalFormat(firstDigits(firstDigitsLength), normalized)
	case 3:
		if numberLength == 9 {
			firstDigitsLength = 2
		}

		if numberLength == 8 {
			return shortFormat(areaCode, normalized)
		}

		return normalFormat(firstDigits(firstDigitsLength), normalized)
	default:
		if numberLength == 9 {
			return shortFormat(areaCode, normalized)
		}

		return normalFormat(firstDigits(2), normalized)
	}
}
