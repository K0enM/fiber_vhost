package fibervhost

import (
	"testing"
)

func Test_Compile_Regexp(t *testing.T) {
	to_match := "test.test.com"
	re, err := string_to_regexp(to_match)
	t.Log(re.String())
	if !(re.String() == "^test.test.com$") {
		t.Error("Error: to_regexp did not compile correctly", re.String())
	} else if err != nil {
		t.Error(err)
	}
}

func Test_String_Wildcard_To_Regexp(t *testing.T) {
	to_test := "*.test.com"
	re, _ := string_to_regexp(to_test)
	to_match := []string{"test.test.com", "yeet.test.com"}
	for _, s := range to_match {
		if !(match(re, s)) {
			t.Error("Error: wildcard did not match any subdomain")
		}
	}
}

func Test_Created_Regexp(t *testing.T) {
	to_re := "^([a-z].example.com)$"
	re, _ := compile_regexp(to_re)
	to_match := "a.example.com"
	not_to_match := "10.example.com"
	if !match(re, to_match) || match(re, not_to_match) {
		t.Error("Error: self created regexp did not work correctly")
	}
}
