package main

import (
	"fmt"
)

func main() {
	fmt.Println(hoge("aa"))
}

func hoge(q string) string {
	return fmt.Sprintf("select * from id = '%s'", q)
}
