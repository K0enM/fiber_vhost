package fibervhost

import (
	"testing"
)

func Test_To_Regexp(t *testing.T) {
	to_test := "yeet"
	re, err := to_regexp(to_test)

	if !(re.String() == "^yeet$") {
		t.Error("Error: to_regexp did not compile correctly", re.String())
	} else if err != nil {
		t.Error(err)
	}
}

func Test_Complete_Regexp(t *testing.T) {
	to_test := "*.yeet.com"
	re := complete_regexp(to_test)
}
