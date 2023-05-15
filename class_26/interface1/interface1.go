package main

import "fmt"

type TapePlayer struct {
    batteries string 
}

func (t TapePlayer) Play(song string) {
    fmt.Println("Playing Using TapePlayer", song)
}

func (t TapePlayer) Stop() {
    fmt.Println("Stopped the TapePlayer")
}



type TapeRecorder struct {
    microphone int
}

func (t TapeRecorder) Play(song string) {
    fmt.Println("Playing using TapeRecorder", song)
}

func (t TapeRecorder) Stop() {
    fmt.Println("Stopped From TapeRecorder")
}

func (t TapeRecorder) Record() {
    fmt.Println("Recording")
}



func Playlist(device TapePlayer, songs []string) {
    for _, song := range songs {
        device.Play(song)
    }
    device.Stop()
}



func main() {
    player := TapePlayer{}
    songs := []string{"a", "b", "c"}
    Playlist(player, songs)

    // recorder := TapeRecorder{}
    // Playlist(recorder, songs)
}

