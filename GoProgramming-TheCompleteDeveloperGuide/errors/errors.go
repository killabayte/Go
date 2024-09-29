package timeparse

import (
	"fmt"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	}
}
