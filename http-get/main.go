package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	contents string
	pos      int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.contents) {
		n := copy(p, []byte(m.contents[m.pos:m.pos+1]))
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	MySlowReaderInstance := &MySlowReader{
		contents: "Hello interfaces!",
	}

	out, err := io.ReadAll(MySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s", out)
}
