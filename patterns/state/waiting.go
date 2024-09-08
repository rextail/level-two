package state

import "fmt"

type WaitingPaymentState struct {
	vending *VendingMachine
}

func NewWaitingPaymentState(vending *VendingMachine) *WaitingPaymentState {
	return &WaitingPaymentState{
		vending: vending,
	}
}

func (w *WaitingPaymentState) ToPaymentConfirmation() {
	fmt.Printf("Вы уже на этапе оплаты")

}

func (w *WaitingPaymentState) CancelPayment() {
	fmt.Println("Возвращаемся к выбору напитков...")
	w.vending.CurrentState = w.vending.decision
}

func (w *WaitingPaymentState) ConfirmPayment(inputSum byte) {
	if inputSum >= w.vending.drinks[w.vending.choice] {
		fmt.Println("Оплата прошла успешно!")
		w.vending.CurrentState = w.vending.giveOut
		return
	}
	diff := w.vending.drinks[w.vending.choice] - inputSum
	fmt.Printf("Введенной суммы недостаточно для оплаты, внесите еще %d", diff)
}

func (w *WaitingPaymentState) GiveOutDrink() {
	fmt.Println("Необходимо дождаться завершения оплаты!")
}
