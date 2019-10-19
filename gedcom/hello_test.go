package gedcom

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, I am a placeholder"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
