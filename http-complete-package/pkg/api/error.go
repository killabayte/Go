package main

type RequestsError struct {
	HTTPCode int
	Body     string
	Err      string
}

func (r RequestsError) Error() string {
	return r.Err
}
