package email

import (
	"testing"
	"fmt"
)

var passingTests = []string {
	"Good@example.com",
	"Good1@example.com",
	"1Good@example.com",
	"3333@example.com",
	"computer@yahoo.com.uk",
	"ba.d@example.com",
}

var failingTests = []string {
	//"computer@.example.com",
	//"computer@example@",
	//"w@example.com",
	//"bad@example@example.com",
	//"bad example@example.com",
	//"bad@example .com",
	"bad@exam.ple.com",
	"bad@example.co.m",
}

var allTests = []string {
	"Good@example.com",
	"Good1@example.com",
	"1Good@example.com",
	"3333@example.com",
	"computer@yahoo.com.uk",
	"ba.d@example.com",
	"computer@.example.com",
	"computer@example@",
	"w@example.com",
	"bad@example@example.com",
	"bad example@example.com",
	"bad@example .com",
	"bad@exam.ple.com",
	"bad@example.co.m",
}

func TestNewAddress(t *testing.T) {

	var result string
	var array1 []string
	for _, test := range allTests {
		value, err := ParseAddress(test)

		if err == nil {
			result = "PASS"
		} else {
			result = "FAIL"
		}
		array1 = append(array1, fmt.Sprintf("(%s) - (%s)\n", value, result))
	}


	fmt.Print("\n\n\n")
	fmt.Println("failing tests are as follows: ")
	fmt.Print(array1, "\n")


	for _, test := range passingTests {
		value, err := ParseAddress(test)
		if err != nil {
			fmt.Println("Failed to parse ", value, err)
		}
	}


	for _, test := range failingTests {
		value, err := ParseAddress(test)
		if err != nil {
			fmt.Println("failed to err email address ", value, err)
		}
	}

}
