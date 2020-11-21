package phone

import (
	"errors"
	"regexp"
	"strings"
)

var phoneRegex = regexp.MustCompile(`\+1[1-9]\d{9}`)
var extractNumbers = regexp.MustCompile(`\d+`)

func ParsePhone(phone string) (string, error) {

	phone = strings.Join(extractNumbers.FindAllString(phone, -1), "")

	if !validatePhone(phone) {
		return phone, errors.New("bad phone number")
	}

	return phone, nil
}

func validatePhone(number string) bool {

	boolean := phoneRegex.MatchString("+" + number)
	if !boolean {
		return false
	}
	return true
}
