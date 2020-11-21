package phone

import (
	"fmt"
	"testing"
)

var passingTests = []string{
	"14089991234",
	"112097772222",
	"19259993344",
	"+1(123) 123-1234",
}

var failingTests = []string{
	"412",
	"",
	"444333111",
	"44443331111",
}

func TestNewPhone(t *testing.T) {

	for _, test := range passingTests {
		value, err := ParsePhone(test)

		if err != nil {
			fmt.Println("Phone number failed to pass: ", value)
			t.Fatal(err)
		}
	}

	for _, test := range failingTests {
		value, err := ParsePhone(test)

		if err == nil {
			fmt.Println("Failed to parse Phone number correctly: ", value)
			t.Fatal(err)
		}
	}
}
