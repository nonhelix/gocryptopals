package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HextoBase64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString([]byte(v)), nil
}
