package email

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	local := "Luke.Ochoa"
	domain := "gmail.com"
	addr, err := NewAddress(local + "@" + domain)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, addr.String(), "Luke.Ochoa@gmail.com")
	fmt.Println(addr)
}
