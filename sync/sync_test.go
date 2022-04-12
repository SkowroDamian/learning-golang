package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup //waitGroup czeka az kolekcja 'goroutines' zakonczy prace. Kazdy goroutine informuje o zakonczeniu pracy
		wg.Add(wantedCount)   // podaje liczbe goroutines na ktore trzeba poczeka

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait() // czeka az wszystkie gouroutines zakocz dziaania

		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func NewCounter() *Counter {
	return &Counter{}
}
