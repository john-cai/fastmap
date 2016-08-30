package lockedval

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkLockedValMap(b *testing.B) {
	lvm := New()
	for i := 0; i < 10; i++ {
		k := strconv.Itoa(i)
		lvm.Set(k, rand.Int())
	}

	wg := &sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			k := strconv.Itoa(rand.Intn(10))
			lvm.Set(k, struct{}{})
			wg.Done()
		}()
	}
	wg.Wait()
}
