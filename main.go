package main

import "github.com/JamesDunne/axewitcher"

func main() {
	// Listen for footswitch activity:
	fswCh, err := ListenFootswitch()
	if err != nil {
		panic(err)
	}

	// Create MIDI interface:
	midi, err := axewitcher.NewMidi()
	if err != nil {
		panic(err)
	}
	defer midi.Close()

	// Initialize controller:
	controller := axewitcher.NewController(midi)

	// Run an idle loop awaiting events:
	for {
		select {
		case ev := <-fswCh:
			controller.HandleFswEvent(ev)
			break
		}
	}
}
