package builder

type PepsiBuilder struct {
	can SodaCan
}

func (p *PepsiBuilder) SetName() {
	p.can.name = "Pepsi-cola"
}
func (p *PepsiBuilder) SetVolume(volume float32) {
	if volume <= 0 {
		panic("can with the given volume can't be created")
	}
	p.can.volume = volume
}
func (p *PepsiBuilder) SetColor() {
	p.can.color = "Blue-red-white"
}
func (p *PepsiBuilder) Build() SodaCan {
	return p.can
}
