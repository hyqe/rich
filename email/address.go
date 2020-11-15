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

var email_regex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ParseAddress(emailString string) (Address, error) {

	if !ValidateAddress(emailString) {
		fmt.Println("\t----This is the bad data: ", emailString)
		return Address{domain: "bad1", local: "data2"}, errors.New("bad data")
	}

	addrSlice := strings.Split(emailString, "@")

	return Address{domain: addrSlice[1], local: addrSlice[0]}, nil
}


func (a Address) String() string {
	return a.local + "@" + a.domain
}

func ValidateAddress(addr string) bool {

	// checks if the address is too short or too long
	if len(addr) < 3 && len(addr) > 254 {
		return false
	}

	// regex checks if the email syntax is valid according to the standard
	boolean := email_regex.MatchString(addr)
	if !boolean  {
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
