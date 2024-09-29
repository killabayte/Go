package timeparse

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}
