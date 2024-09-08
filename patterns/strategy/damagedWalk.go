package strategy

import "fmt"

type DamagedWalk struct {
	velocity byte
	dialogue string
}

func NewDamagedWalk(dialogue string, velocity byte) *DamagedWalk {
	return &DamagedWalk{
		velocity: velocity,
		dialogue: dialogue,
	}
}

func (d *DamagedWalk) Walk() {
	fmt.Println(d.dialogue)
}
