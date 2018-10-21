package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestInsertion_Sort(t *testing.T) {
	sort := Insertion{}
	result := sort.Sort([]int{2, 1, 2})

	t.Log(result)
	t.Fail()
}

func BenchmarkInsertion_Sort(b *testing.B) {
	N := 10
	rand.Seed(time.Now().UnixNano())
	in := make([]int, N)
	for index := 0; index < N; index++ {
		in[index] = rand.Intn(100)
	}
	sort := Insertion{}

	b.StartTimer()
	sort.Sort(in)
	b.StopTimer()

	b.Log(in)
	b.Log(b.N)
}
