package state

import "fmt"

type GiveOutState struct {
	vending *VendingMachine
}

func NewGiveOutState(v *VendingMachine) *GiveOutState {
	return &GiveOutState{v}
}

func (g *GiveOutState) ToPaymentConfirmation() {
	fmt.Println("На данном этапе данная операция невозможна")
}
func (g *GiveOutState) CancelPayment() {
	fmt.Println("Дождитесь получения напитка, он уже оплачен")
}
func (g *GiveOutState) ConfirmPayment(inputSum byte) {
	fmt.Println("Напиток уже оплачен")
}
func (g *GiveOutState) GiveOutDrink() {
	fmt.Printf("Отдаю напиток номер %d\n", g.vending.drinks[g.vending.choice])
	g.vending.CurrentState = g.vending.decision
}
