package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := md5.New()
	hash.Write([]byte("password"))
	a := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(hoge(a))
}

func hoge(q string) string {
	return fmt.Sprintf("select * from id = '%s'", q)
}
