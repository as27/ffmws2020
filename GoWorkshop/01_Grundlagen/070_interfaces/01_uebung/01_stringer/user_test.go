package user

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	u := User{
		"Max",
		"Musterstadt",
		"Musterland",
		66,
	}
	expect := `Max (66)
Musterstadt
Musterland
----------`
	got := fmt.Sprint(u)
	if got != expect {
		t.Errorf("got: \n%s\nexpect: \n%s", u, expect)
	}
}
