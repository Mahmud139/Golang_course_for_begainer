package main

import (
	"golang_tutorial/class_26/gadget"
)

//fixing our playlist function with an interface

type Player interface {
	Play(string)
	Stop()
}

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main(){
	player := gadget.TapePlayer{}
	mixtape := []string{"ami ami", "hello hello", "31 er dingolo"}
	playList(player, mixtape)

	// player := gadget.TapeRecorder{}
	// mixtape := []string{"ami ami", "hello hello", "31 er dingolo"}
	// playList(player, mixtape)
}