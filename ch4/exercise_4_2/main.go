package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hashAlgorithm = flag.String("algo", "sha256", "SHA hash algorithm (sha256, sha384, sha512)")

func main() {
	flag.Parse()
	if *hashAlgorithm != "sha256" && *hashAlgorithm != "sha384" && *hashAlgorithm != "sha512" {
		flag.Usage()
		os.Exit(2)
	}
	stdinReader := bufio.NewScanner(os.Stdin)
	if stdinReader.Scan() {
		input := stdinReader.Text()
		switch *hashAlgorithm {
		case "sha256":
			fmt.Printf("SHA256: %x\n", sha256.Sum256([]byte(input)))
		case "sha384":
			fmt.Printf("SHA384: %x\n", sha512.Sum384([]byte(input)))
		case "sha512":
			fmt.Printf("SHA512: %x\n", sha512.Sum512([]byte(input)))
		}
	}
}
