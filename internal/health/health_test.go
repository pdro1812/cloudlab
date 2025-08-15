package health

import "testing"

func TestOK(t *testing.T) {
	if !OK() {
		t.Fatal("expected true")
	}
}
