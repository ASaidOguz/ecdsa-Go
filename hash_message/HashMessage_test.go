package hashmessage

import (
	"encoding/hex"
	"reflect"
	"testing"

	"golang.org/x/crypto/sha3"
)

func TestHashMessage(t *testing.T) {
	s := "hello world"

	hasKeccak := sha3.NewLegacyKeccak256()
	hasKeccak.Write([]byte(s))
	buf := hasKeccak.Sum(nil)
	want := hex.EncodeToString(buf)

	got := HashMessage(s)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %s want %s ", got, want)
	}
}
