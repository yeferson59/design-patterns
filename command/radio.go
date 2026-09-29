package command

import "fmt"

type Radio struct {
	isRunning bool
}

func NewRadio() *Radio {
	return new(Radio)
}

func (r *Radio) on() {
	r.isRunning = true

	fmt.Println("Radio is on")
}

func (r *Radio) off() {
	r.isRunning = false

	fmt.Println("Radio is off")
}
