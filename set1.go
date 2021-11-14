package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"math"
	"math/bits"
	"unicode/utf8"
)

func hextoBase64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString([]byte(v)), nil
}
func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor:mismatched lengths")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}
func buildCorpus(text string) map[rune]float64 {
	c := make(map[rune]float64)
	for _, char := range text {
		c[char]++
	}
	total := utf8.RuneCountInString(text)
	for char := range c {
		c[char] = c[char] / float64(total)

	}
	return c
}

func scoreEnglish(text string, c map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += c[char]
	}
	return score / float64(utf8.RuneCountInString(text))
}
func singleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, c := range in {
		res[i] = c ^ key
	}
	return res
}
func findSingleXORKey(in []byte, c map[rune]float64) ([]byte, float64) {
	var res []byte
	var bestScore float64
	for key := 0; key < 256; key++ {
		out := singleXOR(in, byte(key))
		score := scoreEnglish(string(out), c)
		if score > bestScore {
			res = out
			bestScore = score
		}
	}
	return res, bestScore
}

func repeatingXOR(in, key []byte) []byte {
	res := make([]byte, len(in))
	for i := range in {
		res[i] = in[i] ^ key[i%len(key)]
	}
	return res
}

func hammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		panic("different length for the two")
	}
	var res int
	for i := range a {
		res += bits.OnesCount8(a[i] ^ b[i])
	}
	return res
}

func findRepeatXORSize(in []byte) int {
	var res int
	var bestScore float64 = math.MaxFloat64
	for keyLen := 2; keyLen < 40; keyLen++ {
		a, b := in[:keyLen*5], in[keyLen*5:keyLen*10]
		score := float64(hammingDistance(a, b)) / float64(keyLen)
		if score < bestScore {
			res = keyLen
			bestScore = score
		}
	}
	return res
}
