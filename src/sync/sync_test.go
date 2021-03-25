package main

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {

	verifyCount := func(t *testing.T, result *Contador, waited int) {
		t.Helper()
		if result.Value() != waited {
			t.Errorf("result %d, waited %d", result.Value(), waited)
		}
	}
	t.Run("Increment the count 3 times result in third value", func(t *testing.T) {
		contador := NewCount()
		contador.Increment()
		contador.Increment()
		contador.Increment()

		verifyCount(t, contador, 3)
	})
	t.Run("run concurrently in safety", func(t *testing.T) {
		countWaited := 1000
		count := NewCount()

		var wg sync.WaitGroup
		wg.Add(countWaited)

		for i := 0; i < countWaited; i++ {
			go func(w *sync.WaitGroup) {
				count.Increment()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verifyCount(t, count, countWaited)
	})
}
