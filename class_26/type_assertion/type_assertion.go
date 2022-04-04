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
	player.Record()
}

func main(){

}