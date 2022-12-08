package sign_message_check_address

import "github.com/ethereum/go-ethereum/crypto"

//SignMessage function takes msg and private_key string and produce ECDSA signature as slice of bytes
func SignMessage(msg string, private_key string) []byte {
	hash := crypto.Keccak256Hash([]byte(msg))
	Private_key, _ := crypto.HexToECDSA(private_key)
	signature, _ := crypto.Sign(hash.Bytes(), Private_key)

	return signature
}
