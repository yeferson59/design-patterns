package state

import (
	"errors"
	"fmt"
)

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func NewItemRequestedState(vendingMachine *VendingMachine) *ItemRequestedState {
	return new(ItemRequestedState{vendingMachine})
}

func (i *ItemRequestedState) requestItem() error {
	return errors.New("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return errors.New("Item Dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}

	fmt.Println("Money entered is ok")

	i.vendingMachine.setState(i.vendingMachine.hasMoney)

	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return errors.New("Please insert money first")
}
