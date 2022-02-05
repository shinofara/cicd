package main

import "testing"

func TestHoge(t *testing.T) {
	h := hoge("aaa")
	if h != "hoge" {
		t.Errorf("want: hoge, got %s", h)
	}
}
