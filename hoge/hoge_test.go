package hoge

import "testing"

func TestHoge(t *testing.T) {
	h := Hoge()
	if h == "" {
		t.Errorf("got %s", h)
	}
}
