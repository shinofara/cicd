package main

import "testing"

func TestHoge(t *testing.T) {
	h := hoge()
	if h != "hoge" {
		t.Errorf("want: hoge, got %s", h)
	}
}
