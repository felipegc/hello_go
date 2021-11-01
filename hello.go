package main

import (
	"fmt"

	"example/user/hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("EPILEF!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
