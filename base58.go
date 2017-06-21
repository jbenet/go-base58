// Copyright (c) 2013-2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
// Modified by Juan Benet (juan@benet.ai)

package base58

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// alphabet is the modified base58 alphabet used by Bitcoin.
const BTCAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const FlickrAlphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode decodes a modified base58 string to a byte slice, using BTCAlphabet
func Decode(b string) ([]byte, error) {
	return DecodeAlphabet(b, BTCAlphabet)
}

// Encode encodes a byte slice to a modified base58 string, using BTCAlphabet
func Encode(b []byte) string {
	return EncodeAlphabet(b, BTCAlphabet)
}

// DecodeAlphabet decodes a modified base58 string to a byte slice, using alphabet.
func DecodeAlphabet(b, alphabet string) ([]byte, error) {
	bigIntVal := big.NewInt(0)
	radix := big.NewInt(58)

	for i := 0; i < len(b); i++ {
		idx := strings.IndexAny(alphabet, string(b[i]))
		if idx == -1 {
			return nil, errors.New("illegal base58 data at input byte " + strconv.FormatInt(int64(i), 10))
		}
		bigIntVal.Mul(bigIntVal, radix)
		bigIntVal.Add(bigIntVal, big.NewInt(int64(idx)))
	}
	temp := bigIntVal.Bytes()

	//append prefix 0
	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabet[0] {
			break
		}
	}
	answerLen := numZeros + len(temp)
	answer := make([]byte, answerLen, answerLen)

	copy(answer[numZeros:], temp)
	return answer, nil
}

// Encode encodes a byte slice to a modified base58 string, using alphabet
func EncodeAlphabet(b []byte, alphabet string) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*137/100+1)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabet[0])
	}

	// reverse
	for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
		answer[i], answer[j] = answer[j], answer[i]
	}

	return string(answer)
}
