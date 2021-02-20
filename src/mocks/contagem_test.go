package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestContagem(t *testing.T) {
	t.Run("print till 3 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Go!`

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("pause before each print", func(t *testing.T) {
		spyPrintSleep := &SpyContagemOperacoes{}
		Contagem(spyPrintSleep, spyPrintSleep)

		waited := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}
		if !reflect.DeepEqual(waited, spyPrintSleep.Chamadas) {
			t.Errorf("waited '%v' calls, result '%v'", waited, spyPrintSleep.Chamadas)
		}
	})

	t.Run("Sleeper configurable", func(t *testing.T) {
		pausedTime := 5 * time.Second

		SpyTime := &TimeSpy{}

		sleeper := SleeperConfigurable{pausedTime, SpyTime.Sleep}
		sleeper.Sleep()
		if SpyTime.durationPause != pausedTime {
			t.Errorf("should paused at %v, but paused at %v", pausedTime, SpyTime.durationPause)
		}
	})
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

// Pausa will pause the Sleep()
func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pause)
}

// Write will put inside the slice the number
func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, write)
	return
}

const write = "escrita"
const pause = "pausa"

type TimeSpy struct {
	durationPause time.Duration
}

func (t *TimeSpy) Sleep(duration time.Duration) {
	t.durationPause = duration
}
