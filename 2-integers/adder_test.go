package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	// to not repeat got != want
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper() // this tells that when a test fails, the failure won't be in here but rather in the caller of this helper
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("2 + 2 = 4", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})

}

// this is another way of writing a test
// provided an example and if the example doesn't return the same value it will fail
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
