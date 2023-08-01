package main

import gadget "github.com/headfirstgo/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
