package phone

import (
	"errors"
	"regexp"
	"strings"
)


var phone_regex = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
var extract_numbers = regexp.MustCompile(`\d+`)

func ParsePhone(phone string) (string, error) {

	phone = strings.Join(extract_numbers.FindAllString(phone, -1), "")

	if !validatePhone(phone) {
		return phone, errors.New("bad phone number")
	}

	return phone, nil
}

func validatePhone(number string) bool {

	boolean := phone_regex.MatchString("+" + number)
	if !boolean {
		return false
	}
	return true
}


































