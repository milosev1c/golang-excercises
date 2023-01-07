package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pf = fmt.Printf

var pc [256]byte

func countBits(sb [sha256.Size]uint8) int {
	var bits int = 0
	for i := range sb {
		bits += int(pc[sb[i]])
	}
	return bits
}

func SHA256BitDiff(sha1, sha2 [sha256.Size]uint8) int {
	b1 := countBits(sha1)
	b2 := countBits(sha2)
	if b1 > b2 {
		return b1 - b2
	}
	return b2 - b1
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	hash1 := sha256.Sum256([]byte(os.Args[1]))
	hash2 := sha256.Sum256([]byte(os.Args[2]))
	for i := int(0); i < len(hash1); i++ {

	}
	_, err := pf("%x\n%x\n%d bits are different\n", hash1, hash2, SHA256BitDiff(hash1, hash2))
	if err != nil {
		return
	}

}
