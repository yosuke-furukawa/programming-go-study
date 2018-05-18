package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var algorithm string

func init() {
	flag.StringVar(&algorithm, "algorithm", "sha256", "hash algorithm like sha256, sha512")
}

func main() {
	flag.Parse()
	stdin := bufio.NewScanner(os.Stdin)
	var c string
	if stdin.Scan() {
		c = stdin.Text()
	}
	algorithm = strings.ToLower(algorithm)
	switch algorithm {
	case "sha256":
		result := sha256.Sum256([]byte(c))
		fmt.Printf("%x\n", result)
	case "sha512":
		result := sha512.Sum512([]byte(c))
		fmt.Printf("%x\n", result)
	case "sha384":
		result := sha512.Sum384([]byte(c))
		fmt.Printf("%x\n", result)
	}

}
