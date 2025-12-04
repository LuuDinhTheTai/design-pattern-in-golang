package adapter

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
