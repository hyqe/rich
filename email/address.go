package email

import (
	"errors"
	"regexp"
	"strings"
)

type Address struct {
	domain string
	local  string
}


var email_regex = regexp.MustCompile(`\b[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\b`)

func ParseAddress(emailString string) (string, error) {

	if !ValidateAddress(emailString) {
		return emailString, errors.New("bad data")
	}

	return emailString, nil
}


func (a Address) String() string {
	return a.local + "@" + a.domain
}

func ValidateAddress(addr string) bool {

	// checks if the address is too short or too long
	if len(addr) < 3 || len(addr) > 254 {
		return false
	}

	// regex checks if the email syntax is valid according to the standard
	boolean := email_regex.MatchString(addr)
	if !boolean  {
		return false
	}

	addrSlice := strings.Split(addr, "@")

	if len(addrSlice[0]) < 3 {
		return false
	}


	return true
}



