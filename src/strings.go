package main

import (
	"fmt"
	"os"
)

func main() {
	s := "hello"
	if s[1] != 'e' {
		os.Exit(1)
	}
	s = "good bye"
	var p *string = &s
	*p = "ciao"

	/*
	以下のコードはstring型の値を変更しようとしているので不正なコード
	s[0] = 'x' // cannot assign to s[0]
	(*p)[1] = 'y' // cannot assign to (*p)[1]
	*/
	fmt.Printf(s)
}
