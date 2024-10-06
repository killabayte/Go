package main

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
