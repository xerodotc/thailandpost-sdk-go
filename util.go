package thailandpost

import (
	"regexp"
)

var upsS10Regexp = regexp.MustCompile(`^([A-Z]{2})([0-9]{9})([A-Z]{2})$`)
var upsS10ChecksumWeight = []int{8, 6, 4, 2, 3, 5, 9, 7}

// Check whether a tracking number is comply with UPS S10 format
func IsValidTrackingNumber(tracking string) bool {
	if !upsS10Regexp.MatchString(tracking) {
		return false
	}

	serial := tracking[2:10]
	checkDigit := int(tracking[10]) - '0'

	sum := 0
	for i, digit := range serial {
		sum += (int(digit) - '0') * upsS10ChecksumWeight[i]
	}

	calcCheckDigit := 11 - (sum % 11)
	if calcCheckDigit == 10 {
		calcCheckDigit = 0
	} else if calcCheckDigit == 11 {
		calcCheckDigit = 5
	}

	if checkDigit != calcCheckDigit {
		return false
	}

	return true
}

// Check multiple tracking numbers
// Return InvalidTrackingNumbersError with list of invalid tracking numbers
func ValidateTrackingNumbers(trackings []string) error {
	invalidTrackings := make([]string, 0, len(trackings))
	for _, tracking := range trackings {
		if !IsValidTrackingNumber(tracking) {
			invalidTrackings = append(invalidTrackings, tracking)
		}
	}
	if len(invalidTrackings) > 0 {
		return InvalidTrackingNumbersError{
			InvalidTrackingNumbers: invalidTrackings,
		}
	}
	return nil
}
