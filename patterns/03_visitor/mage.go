package visitor

type Mage struct {
	name string
}

func NewMage() *Mage {
	return &Mage{
		name: "Mage",
	}
}

func (m *Mage) accept(v Visitor) {
	v.VisitForMage(m)
}
