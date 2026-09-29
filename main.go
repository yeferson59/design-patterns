package main

import (
	"fmt"
	"log"

	"github.com/yeferson59/design-patterns/command"
	"github.com/yeferson59/design-patterns/state"
)

func main() {
	fmt.Println("Patrón de Diseño Command")
	fmt.Println()

	tv := command.NewTv()

	onCommand, offCommand := command.NewOnCommand(tv), command.NewOffCommand(tv)

	onButton, offButton := command.NewButton(onCommand), command.NewButton(offCommand)

	onButton.Press()
	offButton.Press()

	radio := command.NewRadio()

	radioOnCommand := command.NewOnCommand(radio)
	radioOffCommand := command.NewOffCommand(radio)

	radioOnButton := command.NewButton(radioOnCommand)
	radioOffButton := command.NewButton(radioOffCommand)

	radioOnButton.Press()
	radioOffButton.Press()

	fmt.Println("\nPatrón de Diseño State")
	fmt.Println()

	vendingMachine := state.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
}
