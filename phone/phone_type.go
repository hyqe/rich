package phone

import (
	"errors"
	"regexp"
)


var phone_regex = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)


func ParsePhone(phone string) (string, error) {

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


































