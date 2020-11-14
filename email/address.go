package email

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

type Address struct {
	domain string
	local  string
}

func ParseAddress(emailString string) (Address, error) {
	// TODO: check validity of local & domain

	if !ValidateAddress(emailString) {
		return Address{domain: "bad1", local: "data2"}, errors.New("bad data")
	}

	addrSlice := strings.Split(emailString, "@")

	return Address{domain: addrSlice[1], local: addrSlice[0]}, nil
}


func (a Address) String() string {
	return a.local + "@" + a.domain
}

func ValidateAddress(addr string) bool {
	fmt.Println("Did you pass by here?")

	// checks if the address is too short or too long
	if len(addr) < 3 && len(addr) > 254 {
		return false
	}

	// regex checks if the email syntax is valid according to the standard
	_, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", "doesthiswork@gmail.com")
	if err != nil {
		return false
	}

	// LookupMX checks if the domain is real
	addrSlice := strings.Split(addr, "@")
	mx, err2 := net.LookupMX(addrSlice[1])
	if err2 != nil || len(mx) == 0 {
		return false
	}

	return true
}
