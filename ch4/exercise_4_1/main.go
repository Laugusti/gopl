// diffSha256Bits counts the number of bits that are different in two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/Laugusti/gopl/ch2/popcount"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage diffSha256Bits STRING1 STRING2")
		os.Exit(1)
	}
	h1 := sha256.Sum256([]byte(os.Args[1]))
	h2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("hash1: %x\nhash2: %x\nbits diff: %d\n", h1, h2, diffSha256Bits(h1, h2))
}

func diffSha256Bits(hash1, hash2 [sha256.Size]byte) int {
	var result int
	for i, byte1 := range hash1 {
		byte2 := hash2[i]
		result += popcount.PopCount(uint64(byte1 ^ byte2))
	}
	return result
}
