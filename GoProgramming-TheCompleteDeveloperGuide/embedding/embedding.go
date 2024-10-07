package main

import "fmt"

type Whisperer interface {
	Whisper() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whisperer
	Yeller
}

func talk(t Talker) {
	fmt.Println(t.Yell())

	fmt.Println(t.Whisper())
}
