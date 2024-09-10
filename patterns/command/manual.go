package command

//Те самые инвокеры

type ManualEnabler struct {
	enabler DeviceEnabler
}

type ManualDisabler struct {
	disabler DeviceDisabler
}

func (d *ManualEnabler) execute() bool {
	return d.enabler.Enable()
}

func (d *ManualDisabler) execute() bool {
	return d.disabler.Disable()
}
