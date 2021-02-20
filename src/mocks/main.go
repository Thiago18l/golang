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

// SleeperConfigurable it's a implemention of time.Sleep
type SleeperConfigurable struct {
	duration time.Duration
	pause    func(time.Duration)
}

// SleeperSpy see the number of 'Chamadas'
type SleeperSpy struct {
	Chamadas int
}

// Sleep will count the time in seconds
func (s *SleeperConfigurable) Sleep() {
	s.pause(s.duration)
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
	sleeper := &SleeperConfigurable{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
