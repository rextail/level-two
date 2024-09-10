package builder

type CocaColaBuilder struct {
	can SodaCan
}

func (c *CocaColaBuilder) SetName() {
	c.can.name = "Coca-cola"
}
func (c *CocaColaBuilder) SetVolume(volume float32) {
	if volume <= 0 {
		panic("can with the given volume can't be created")
	}
	c.can.volume = volume
}
func (c *CocaColaBuilder) SetColor() {
	c.can.color = "red-white"
}
func (c *CocaColaBuilder) Build() SodaCan {
	return c.can
}
