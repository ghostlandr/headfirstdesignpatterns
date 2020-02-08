package objects

import "fmt"

type Stereo struct {}

func (s *Stereo) On() {
	fmt.Println("Turning stereo on")
}

func (s *Stereo) Off() {
	fmt.Println("Turning stereo off")
}

func (s *Stereo) SetCD() {
	fmt.Println("Playing CD")
}

func (s *Stereo) SetDVD() {
	fmt.Println("Playing DVD")
}

func (s *Stereo) SetRadio() {
	fmt.Println("Playing radio")
}

func (s *Stereo) SetVolume(level int) {
	fmt.Printf("Setting volume to %d\n", level)
}

