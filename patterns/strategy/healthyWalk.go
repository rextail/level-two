package strategy

import "fmt"

type HealthyWalk struct {
	velocity byte
	dialogue string
}

func NewHealthyWalk(dialogue string, velocity byte) *HealthyWalk {
	return &HealthyWalk{
		velocity: velocity,
		dialogue: dialogue,
	}
}
func (h *HealthyWalk) Walk() {
	fmt.Println(h.dialogue)
}
