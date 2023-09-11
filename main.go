package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-varint"
)

// BTCAlphabet is the bitcoin base58 alphabet.
var BTCAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var (
	bn58  = big.NewInt(58)
	bn256 = big.NewInt(256)
)

func ForgePeerID(s string) (peer.ID, error) {
	remainder := big.NewInt(0)

	for i, c := range s {
		index := strings.IndexRune(BTCAlphabet, c)
		if index == -1 {
			return "", fmt.Errorf("invalid character %q: not in base58 alphabet", c)
		}
		p58 := new(big.Int).Exp(bn58, big.NewInt(int64(len(s)-i-1)), new(big.Int))
		remainder.Add(remainder, new(big.Int).Mul(p58, big.NewInt(int64(index))))
	}
	base0 := big.NewInt(256)
	for base0.Cmp(remainder) < 0 {
		base0.Mul(base0, big.NewInt(256))
	}

	length := uint64(len(s)*406/555 + 1) // approximate of ceil(log(256)/log(58))
	modulus := new(big.Int).Exp(bn58, big.NewInt(int64(len(s))), new(big.Int))

	prefix := append([]byte{0}, varint.ToUvarint(length)...)

	base := new(big.Int).Mul(new(big.Int).SetBytes(prefix), new(big.Int).Exp(bn256, big.NewInt(int64(length)), new(big.Int)))

	_, r := new(big.Int).DivMod(base, modulus, new(big.Int))

	base.Add(base, remainder)
	base.Sub(base, r)

	if r.Cmp(remainder) > 0 {
		base.Add(base, modulus)
	}

	p, err := peer.IDFromBytes(append([]byte{0}, base.Bytes()...))
	if err != nil {
		panic(err)
	}
	return p, nil
}

func main() {
	if len(os.Args) != 2 || os.Args[1][0] == '-' {
		fmt.Println("  Usage: ./forge <base58 peerid suffix>\n  Example: ./forge ooooPEER\n          > 1EooooPEER")
		os.Exit(0)
	}
	forged, err := ForgePeerID(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	fmt.Println(forged)
}
