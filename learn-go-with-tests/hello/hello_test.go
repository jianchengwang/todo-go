package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Dog", "english")
		want := "Hello Dog"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to empty string", func(t *testing.T) {
		got := Hello("Dog", "french")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
