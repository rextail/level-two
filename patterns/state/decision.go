package state

import "fmt"

// DecisionState Состояние, в котором можно тыкать кнопочки и передумывать без проблем
type DecisionState struct {
	vending *VendingMachine
}

func NewDecisionState(v *VendingMachine) *DecisionState {
	return &DecisionState{v}
}

func (d *DecisionState) ToPaymentConfirmation() {
	fmt.Printf("Переходим к этапу оплаты, сумма к оплате: %d\n", d.vending.drinks[d.vending.choice])
	d.vending.CurrentState = d.vending.waiting
}

func (d *DecisionState) CancelPayment() {
	fmt.Println("Вы еще не дошли до этапы оплаты")
}

func (d *DecisionState) ConfirmPayment(inputSum byte) {
	fmt.Println("Видимо, какой-то сбой, мы не ждали оплату")
}

func (d *DecisionState) GiveOutDrink() {
	fmt.Println("Видимо, какой-то сбой, сейчас невозможно отдать напиток")
}
