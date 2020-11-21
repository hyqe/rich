package phone


import (
	"fmt"
	"testing"
)



//var passingTests []string {
//	"4089991234",
//	"2097772222",
//	"9259993344",
//}
//
//var failingTests []string {
//	"412",
//	"",
//	"444333111",
//	"44443331111",
//}


var allTests = []string {
	"412",
	"",
	"444333111",
	"44443331111",
	"4089991234",
	"2097772222",
	"9259993344",
	"+1(123) 123-1234",
}

func TestNewPhone(t *testing.T) {

	var array1 []string
	var result string
	for _, test := range allTests {
		value, err := ParsePhone(test)

		if err == nil {
			result = "PASS"
		} else {
			result = "FAIL"
		}
		array1 = append(array1, fmt.Sprintf("(%s) - (%s)\n", value, result))
	}

	fmt.Println(array1)






}
