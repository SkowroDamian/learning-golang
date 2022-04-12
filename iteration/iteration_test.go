package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("testing for A", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := "aaaa"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

	t.Run("testing for D", func(t *testing.T) {
		repeated := Repeat("d", 6)
		expected := "dddddd"

		if repeated != expected {
			t.Errorf("Expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
