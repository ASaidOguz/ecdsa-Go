package findfavoritecolor

import (
	"crypto/sha256"
	"reflect"
)

//Findcolor function find and returns the correct hashed string
func Findcolor(hash []byte) string {
	Colors := [6]string{"red", "green", "blue", "yellow", "pink", "orange"}
	message := "Theres no matching!"
	for _, color := range Colors {
		h := sha256.New()
		h.Write([]byte(color))
		bs := h.Sum(nil)
		isValid := reflect.DeepEqual(bs, hash)
		if isValid == true {
			return color
		}
	}
	return message
}
