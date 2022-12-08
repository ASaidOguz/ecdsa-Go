package findfavoritecolor

import (
	"crypto/sha256"
	"testing"
)

func TestFindcolor(t *testing.T) {

	s := "red"
	h := sha256.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	got := Findcolor(bs)
	want := "red"
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
