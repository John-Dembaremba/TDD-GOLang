package mocking

import (
	"fmt"
	"io"
)

const (
	countStart = 3
	finalWord  = "Go!"
)

type Sleeper interface {
	Sleep()
}

func CountDown(writer io.Writer, sleeper Sleeper) {
	
	for v := countStart; v > 0; v-- {
		fmt.Fprintln(writer, v)
		sleeper.Sleep()
	}
	fmt.Fprintln(writer, finalWord)
}
