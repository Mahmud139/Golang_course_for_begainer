package main

import (
	"golang_tutorial/class_26/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("mahmud")
	player.Stop()
	var recorder gadget.TapeRecorder = player.(gadget.TapeRecorder)
	recorder.Record()
}

func main(){
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}