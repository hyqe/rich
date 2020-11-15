package email

import (
	"fmt"
	"testing"
)

func TestNewAddress(t *testing.T) {
	local := "Luke.Ochoa"
	domain := "gmail.com"
	addr, err := ParseAddress(local + "@" + domain)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(addr)

	addr, err = ParseAddress(" BadEmail" + "@" + "fake.lol")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addr)


}
