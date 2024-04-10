package example

import (
	"os"
	"regexp"
)

func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

func FindPhoneNumber(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

type Repo struct {
	Value int
}

func NewRepo(v int) *Repo {
	return &Repo{Value: v}
}
