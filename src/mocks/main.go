package mocks

import (
	"fmt"
	"io"
	"os"
	"time"
)

const lastWord = "Go!"
const index = 3

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// SleeperSpy see the number of 'Chamadas'
type SleeperSpy struct {
	Chamadas int
}

// SleeperDefault do nothing at this time
type SleeperDefault struct{}

// Sleep will add a count to 'Chamadas'
func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

// Sleep will count the time in seconds
func (d *SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

// Contagem will print a output in screen
func Contagem(exit io.Writer, sleeper Sleeper) {
	for i := index; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(exit, i)
	}
	sleeper.Sleep()
	fmt.Fprint(exit, lastWord)
}

func main() {
	sleeper := &SleeperDefault{}
	Contagem(os.Stdout, sleeper)
}
