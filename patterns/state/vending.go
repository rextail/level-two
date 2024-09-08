package state

import "fmt"

const (
	soda = byte(iota)
	juice
	water
)

type VendingMachine struct {
	waiting      *WaitingPaymentState
	decision     *DecisionState
	giveOut      *GiveOutState
	CurrentState State
	drinks       map[byte]byte //drink id and it's cost
	choice       byte
}

func NewVendingMachine() *VendingMachine {
	vending := &VendingMachine{drinks: map[byte]byte{soda: 50, juice: 65, water: 24}}
	vending.decision = NewDecisionState(vending)
	vending.waiting = NewWaitingPaymentState(vending)
	vending.giveOut = NewGiveOutState(vending)
	return vending
}

func (v *VendingMachine) SelectDrink(id byte) {
	if _, ok := v.drinks[id]; !ok {
		fmt.Println("Странно, у нас нет такого напитка")
		return
	}
	v.choice = id
	fmt.Println("Хороший выбор!")
}

func (v *VendingMachine) UpdatePrice(id byte, newPrice byte) bool {
	if _, ok := v.drinks[id]; !ok {
		return false
	}
	v.drinks[id] = newPrice
	return true
}
