package sign_message_check_address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func Test_signMessage(t *testing.T) {

	private_key := "c4429120d243ec311ae290d29530672ffaccb6dcb7e571a7f51f9398d5532fa4"
	msg := "hello"
	hash := crypto.Keccak256Hash([]byte(msg))

	Private_key, _ := crypto.HexToECDSA(private_key)
	Public_key := Private_key.PublicKey

	publicKeyBytes := crypto.FromECDSAPub(&Public_key)
	fmt.Println("Public key: ", hex.EncodeToString(publicKeyBytes))
	signature := SignMessage(msg, private_key)
	sigPubkey, _ := crypto.Ecrecover(hash[:], signature)
	address := crypto.PubkeyToAddress(Public_key)
	fmt.Println("------------------------------->")
	fmt.Println("Address derived from Public key: ", address)
	fmt.Println("------------------------------->")
	if bytes.Equal(sigPubkey, publicKeyBytes) {
		fmt.Println("MATCHING COMPLETED!!")
	} else {
		fmt.Println("MATCHING FAILED!!")
	}
}
