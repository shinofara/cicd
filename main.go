package main

import (
	"fmt"
	"io"

	"golang.org/x/crypto/md4"
)

func main() {
	fmt.Println(hoge())
	h := md4.New()
	data := "These pretzels are making me thirsty."
	io.WriteString(h, data)
	fmt.Printf("MD4 is the new MD5: %x\n", h.Sum(nil))
}

func hoge() string {
	return "hoge"
}
