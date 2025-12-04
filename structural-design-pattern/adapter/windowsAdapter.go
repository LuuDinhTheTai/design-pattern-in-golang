package adapter

import "fmt"

type WindowAdapter struct {
	windowMachine *Windows
}

func (w *WindowAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
