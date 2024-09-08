package visitor

type Druid struct {
	name string
}

func NewDruid() *Druid {
	return &Druid{
		name: "Druid",
	}
}

func (d *Druid) accept(v Visitor) {
	v.VisitForDruid(d)
}
