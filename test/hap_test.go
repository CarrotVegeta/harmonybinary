package test

import (
	"fmt"
	"harmonybinary/hap"
	"testing"
)

func TestHap(t *testing.T) {
	a, err := hap.OpenFile("./c.hap")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
