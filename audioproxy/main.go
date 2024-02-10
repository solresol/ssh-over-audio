package main

import "fmt"
import "github.com/gordonklaus/portaudio"

func main() {
    portaudio.Initialize()

    defer portaudio.Terminate()
    go microphoneReader()
    go audioWriter()
    fmt.Println("audioproxy placeholder")
}

// microphoneReader reads 100ms of audio samples and puts them into an array
func microphoneReader() {
    // TODO: Use PortAudio to read from the microphone
}

// audioWriter creates 50ms of sine wave data in an array and writes it out
func audioWriter() {
    // TODO: Use PortAudio to write to the audio output device
}
