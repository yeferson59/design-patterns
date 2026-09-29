package state

import (
	"errors"
)

type NoItemState struct {
	vendingMachine *VendingMachine
}

func NewNoItemState(vendingMachine *VendingMachine) *NoItemState {
	return new(NoItemState{vendingMachine})
}

func (i *NoItemState) requestItem() error {
	return errors.New("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)

	i.vendingMachine.setState(i.vendingMachine.hasItem)

	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return errors.New("Item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return errors.New("Item out of stock")
}
