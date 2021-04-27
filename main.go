package main

import (
	"fmt"

	"github.com/hsmtkk/urban-invention/beersong"
)

func main() {
	for n := 99; n >= 0; n-- {
		phrase := beersong.GetPhrase(n)
		fmt.Println(phrase)
	}
}
