package fibervhost

import (
	"regexp"
	"strings"
)

const AsteriskRegexp = "/\\*/g"
const AsteriskReplace = "([^.]+)"
const EndAnchoredRegexp = "/(?:^|[^\\])(?:\\\\)*\\$$/"
const EscapeRegexp = "/([.+?^=!:${}()|[\\]/\\])/g"
const EscapeReplace = "\\$1"

func compile_regexp(to_match string) (*regexp.Regexp, error) {
	to_match = enforce_start_end_characters(to_match)
	return regexp.Compile(to_match)
}

func to_regexp(string to_match) (*regexp.Regexp, error) {
	strings.Replace(to_match, EscapeRegexp, EscapeReplace, 1)
	strings.Replace(to_match, AsteriskRegexp, AsteriskReplace, -1)
	return compile_regexp(str)
}

func complete_regexp(re *regexp.Regexp) (*regexp.Regexp, error) {
	str := re.String()
	return compile_regexp(str)
}

func enforce_start_end_characters(str string) string {
	if string(str[0]) != string("^") {
		str = string("^") + str
	}

	anchor_test := regexp.Compile(EndAnchoredRegexp)

	if !anchor_test.MatchString(str) {
		str = str + string("$")
	}
	return str
}
