package strategy

import "fmt"

type InjuredWalk struct {
	velocity byte
	dialogue string
}

func NewInjuredWalk(dialogue string, velocity byte) *InjuredWalk {
	return &InjuredWalk{
		velocity: velocity,
		dialogue: dialogue,
	}
}

func (i *InjuredWalk) Walk() {
	fmt.Println(i.dialogue)
}
