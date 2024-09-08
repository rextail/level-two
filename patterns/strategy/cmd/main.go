package main

import "strategy"

func main() {
	//создадим нового персонажа
	char := strategy.NewCharacter()
	//обновим алгоритм передвижения
	char.UpdateWalkingStrategy()
	char.DoWalk()
	//поцарапаем персонажа, чтобы изменить его ходьбу
	char.HP -= 30
	char.UpdateWalkingStrategy()
	char.DoWalk()
	//травмируем персонажа, чтобы изменить его ходьбу
	char.HP -= 50
	char.UpdateWalkingStrategy()
	char.DoWalk()
}
