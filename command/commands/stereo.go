package commands

import (
	"headfirstdesigntraining/command/objects"
)

func NewStereoOnWithCD(s *objects.Stereo) Command {
	return &stereoOnWithCDCommand{s: s}
}

type stereoOnWithCDCommand struct {
	s *objects.Stereo
}

func (s *stereoOnWithCDCommand) Execute() {
	s.s.On()
	s.s.SetCD()
	s.s.SetVolume(11)
}

func (s *stereoOnWithCDCommand) Undo() {
	s.s.Off()
}

func (s *stereoOnWithCDCommand) Name() string {
	return "Turn on the stereo in CD mode, volume @ 11"
}

func NewStereoOff(s *objects.Stereo) Command {
	return &stereoOff{s: s}
}

type stereoOff struct {
	s *objects.Stereo
}

func (s *stereoOff) Execute() {
	s.s.Off()
}

func (s *stereoOff) Undo() {
	s.s.On()
}

func (s *stereoOff) Name() string {
	return "Turn off stereo"
}

