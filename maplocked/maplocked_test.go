package maplocked

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkLockedMap(b *testing.B) {
	lm := New()
	for i := 0; i < 10; i++ {
		k := strconv.Itoa(i)
		lm.Set(k, rand.Int())
	}

	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			k := strconv.Itoa(rand.Intn(10))
			lm.Set(k, struct{}{})
			wg.Done()
		}()
	}
	wg.Wait()
}
