package email

import (
	"fmt"
	"testing"
)

var passingTests = []string{
	"Good@example.com",
	"Good1@example.com",
	"1Good@example.com",
	"3333@example.com",
	"computer@yahoo.com.uk",
	"ba.d@example.com",
	"bad@exam.ple.com",
}

var failingTests = []string{
	"computer@.example.com",
	"computer@example@",
	"w@example.com",
	"bad@example@example.com",
	"bad example@example.com",
	"bad@example .com",
	"bad@example.co.m",
}

func TestNewAddress(t *testing.T) {

	for _, test := range passingTests {
		value, err := ParseAddress(test)
		if err != nil {
			fmt.Println("Failed to parse ", value)
			t.Fatal(err)
		}
	}

	for _, test := range failingTests {
		value, err := ParseAddress(test)
		if err == nil {
			fmt.Println("failed to err email address ", value)
			t.Fatal(err)
		}
	}

}
