package command

import "fmt"

type Tv struct {
	isRunning bool
}

func NewTv() *Tv {
	return new(Tv)
}

func (t *Tv) on() {
	t.isRunning = true

	fmt.Println("TV is on")
}

func (t *Tv) off() {
	t.isRunning = false

	fmt.Println("TV is off")
}
