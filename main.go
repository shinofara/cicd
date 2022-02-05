package main

import "fmt"

func main() {
	fmt.Println(hoge("'a'; Delete from user"))
}

func hoge(q string) string {
	return fmt.Sprintf("select * from id = %s", q)
}
