package main

import (
	"fmt"
	"headfirstdesigntraining/command/remote"
	"headfirstdesigntraining/command/objects"
	"headfirstdesigntraining/command/commands"
)

func main() {
	r := remote.New()

	// Add light switches
	l := &objects.Light{}
	lonc := commands.NewLightOn(l)
	loffc := commands.NewLightOff(l)
	r.SetCommand(0, lonc, loffc)

	// Add stereo control
	s := &objects.Stereo{}
	sOn := commands.NewStereoOnWithCD(s)
	sOff := commands.NewStereoOff(s)
	r.SetCommand(1, sOn, sOff)

	// Add party mode macro
	mac := commands.NewMacro("Party Mode", []commands.Command{lonc, sOn})
	r.SetCommand(2, mac, commands.NewNoop())

	fmt.Println(r)

//	r.OnButtonWasPressed(0)
//	fmt.Println(r)
//	r.OnButtonWasPressed(1)
//	fmt.Println(r)
//	r.OffButtonWasPressed(0)
//	fmt.Println(r)
//	r.UndoButtonWasPressed()
//	r.OffButtonWasPressed(1)

	r.OnButtonWasPressed(2)
	r.UndoButtonWasPressed()
}
