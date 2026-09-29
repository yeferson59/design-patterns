package state

import (
	"errors"
	"fmt"
)

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func NewHasMoneyState(vendingMachine *VendingMachine) *HasMoneyState {
	return new(HasMoneyState{vendingMachine})
}

func (i *HasMoneyState) requestItem() error {
	return errors.New("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return errors.New("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return errors.New("Item out of stock")
}

func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")

	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1

	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}

	return nil
}
