package integers

import "testing"

var (
	fixedSize   = []int{1, 2, 3, 4, 5}
	sliceValues = []int{1, 2, 3}
)

func TestAdd(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		got := Add(fixedSize)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, fixedSize)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		got := Add(sliceValues)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, sliceValues)
		}
	})
}

func BenchmarkAdd(t *testing.B) {
	t.Run("collection of 5 numbers", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			Add(fixedSize)
		}
	})
	t.Run("collection of any size", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			Add(sliceValues)
		}
	})
}
