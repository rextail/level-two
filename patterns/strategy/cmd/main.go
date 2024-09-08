package main

import "strategy"

const (
	damagedSpeed = 25
	injuredSpeed = 50
	healthySpeed = 100
)

var healthy = strategy.NewHealthyWalk("la-la-la", healthySpeed)
var injured = strategy.NewInjuredWalk("oh.. hh..", injuredSpeed)
var damaged = strategy.NewDamagedWalk("h-h-h-elp, s-some-b-b-ody", damagedSpeed)

func main() {
	//создадим нового персонажа
	char := strategy.NewCharacter()

	char.SetWalker(healthy)
	char.DoWalk()

	//поцарапаем персонажа, чтобы изменить его ходьбу
	char.HP -= 30
	if char.HP > 50 && char.HP < 75 {
		char.SetWalker(injured)
	}
	char.DoWalk()

	//травмируем персонажа, чтобы изменить его ходьбу
	char.HP -= 50
	if char.HP > 1 && char.HP <= 25 {
		char.SetWalker(damaged)
	}
	char.DoWalk()
}
