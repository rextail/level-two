package main

import "strategy"

func main() {
	healthy := strategy.NewHealthyWalk("la-la-la", 100)
	injured := strategy.NewDamagedWalk("oh.. hh..", 50)
	damaged := strategy.NewDamagedWalk("h-h-h-elp, s-some-b-b-ody", 25)
	char := strategy.NewCharacter()
	if char.HP >= 75 {
		char.Walk = healthy
		char.DoWalk()
		char.HP -= 26
	}
	if char.HP >= 25 && char.HP < 75 {
		char.Walk = injured
		char.DoWalk()
		char.HP -= 60
	}
	if char.HP > 1 && char.HP < 25 {
		char.Walk = damaged
		char.DoWalk()
		char.Walk = damaged
	}
}
