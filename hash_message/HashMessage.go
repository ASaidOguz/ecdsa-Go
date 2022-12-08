package hashmessage

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

//HashMessage taking string and hash it with Keccak256 hash(legacy)
func HashMessage(msg string) string {
	hashKeccak := sha3.NewLegacyKeccak256()
	hashKeccak.Write([]byte(msg))
	buf := hashKeccak.Sum(nil)
	return hex.EncodeToString(buf)
}
