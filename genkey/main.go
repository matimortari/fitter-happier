package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

const defaultBytes = 32

func main() {
	n := defaultBytes

	if len(os.Args) > 1 {
		parsed, err := strconv.Atoi(os.Args[1])
		if err != nil || parsed <= 0 {
			fmt.Fprintln(os.Stderr, "usage: genkey [bytes]")
			os.Exit(1)
		}
		n = parsed
	}

	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		fmt.Fprintln(os.Stderr, "error generating random bytes:", err)
		os.Exit(1)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(buf))
}
