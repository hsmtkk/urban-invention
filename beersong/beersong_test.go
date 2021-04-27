package beersong_test

import (
	"testing"

	"github.com/hsmtkk/urban-invention/beersong"
	"github.com/stretchr/testify/assert"
)

func TestGetPhrase3(t *testing.T) {
	want := `3 bottles of beer on the wall, 3 bottles of beer.
Take one down and pass it around, 2 bottles of beer on the wall.`
	got := beersong.GetPhrase(3)
	assert.Equal(t, want, got)
}

func TestGetPhrase2(t *testing.T) {
	want := `2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.`
	got := beersong.GetPhrase(2)
	assert.Equal(t, want, got)
}

func TestGetPhrase1(t *testing.T) {
	want := `1 bottle of beer on the wall, 1 bottle of beer.
Take it down and pass it around, no more bottles of beer on the wall.`
	got := beersong.GetPhrase(1)
	assert.Equal(t, want, got)
}

func TestGetPhrase0(t *testing.T) {
	want := `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.`
	got := beersong.GetPhrase(0)
	assert.Equal(t, want, got)
}
