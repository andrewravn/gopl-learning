package main

import "testing"

func TestRepeat(t *testing.T) {
	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("Repeat a 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
