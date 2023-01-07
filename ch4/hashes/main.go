package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	hashType := flag.String("t", "sha256", "Тип хеша")
	flag.Parse()
	switch *hashType {
	case "sha256":
		fmt.Printf("%x", sha256.Sum256([]byte(os.Args[1])))
	case "sha384":
		fmt.Printf("%x", sha512.Sum384([]byte(os.Args[1])))
	case "sha512":
		fmt.Printf("%x", sha512.Sum512([]byte(os.Args[1])))
	}

}
