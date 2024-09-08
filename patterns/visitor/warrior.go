package visitor

type Warrior struct {
	name string
}

func NewWarrior() *Warrior {
	return &Warrior{
		name: "Warrior",
	}
}
func (w *Warrior) accept(v Visitor) {
	v.VisitForWarrior(w)
}
