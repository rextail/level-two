package strategy

//Поведенческий паттерн, который позволяет менять поведение объекта в рантайме

//Применяется, когда объект предполагает изменение поведения в зависимости от какого-то условия.
//Согласно идее паттерна, мы создаем семейство алгоритмов, где каждый алгоритм будет иметь собственный класс,
//который удовлетворяет общему интерфейсу. Таким образом добиваемся взаимозаменяемости алгоритмов.

// Например, у нас есть персонаж в компьютерной игре:
// Когда персонаж имеет достаточно здоровья, у него обычная анимация ходьбы.
// Когда персонаж немного просаживается по здоровью, он начинает хромать, скорость медленная
// Когда персонаж значительно просаживается по здоровью, он начинает ползти, скорость очень медленная
const (
	damagedSpeed = 25
	injuredSpeed = 50
	healthySpeed = 100
)

var healthy = NewHealthyWalk("la-la-la", healthySpeed)
var injured = NewDamagedWalk("oh.. hh..", injuredSpeed)
var damaged = NewDamagedWalk("h-h-h-elp, s-some-b-b-ody", damagedSpeed)

type Walker interface {
	Walk()
}

type Character struct {
	HP   byte
	Walk Walker //🐺🐺🐺
}

func NewCharacter() *Character {
	return &Character{
		HP: 100,
	}
}

func (c *Character) DoWalk() {
	c.Walk.Walk()
}

func (c *Character) UpdateWalkingStrategy() {
	if c.HP >= 75 {
		c.Walk = healthy

	}
	if c.HP >= 25 && c.HP < 75 {
		c.Walk = injured
	}
	if c.HP > 1 && c.HP < 25 {
		c.Walk = damaged
	}
}
