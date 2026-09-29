package state

import (
	"errors"
	"fmt"
)

type HasItemState struct {
	vendingMachine *VendingMachine
}

func NewHasItemState(vendingMachine *VendingMachine) *HasItemState {
	return new(HasItemState{vendingMachine})
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)

		return errors.New("No item present")
	}

	fmt.Println("Item requestd")

	i.vendingMachine.setState(i.vendingMachine.itemRequested)

	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)

	i.vendingMachine.IncrementItemCount(count)

	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return errors.New("Please select item first")
}

func (i *HasItemState) dispenseItem() error {
	return errors.New("Please select item first")
}
