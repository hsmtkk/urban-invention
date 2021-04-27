package beersong

import "fmt"

func GetPhrase(n int) string {
	if n == 0 {
		return `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.`
	}
	if n == 1 {
		return `1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.`
	}
	if n == 2 {
		return `2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.`
	}
	return fmt.Sprintf(`%d bottles of beer on the wall, %d bottles of beer.
Take one down and pass it around, %d bottles of beer on the wall.`, n, n, n-1)
}
