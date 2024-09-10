package command

//Те самые инвокеры

type DistantEnabler struct {
	enabler DeviceEnabler
}

type DistantDisabler struct {
	disabler DeviceDisabler
}

func (d *DistantEnabler) execute() bool {
	return d.enabler.Enable()
}

func (d *DistantDisabler) execute() bool {
	return d.disabler.Disable()
}
