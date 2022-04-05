package main

import (
	"fmt"
	"golang_tutorial/class_26/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("mahmud")
	player.Stop()
	//var recorder gadget.TapeRecorder = player.(gadget.TapeRecorder)
	//recorder.Record()

	//avoiding type assertion failure
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
	
}

func main(){
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}