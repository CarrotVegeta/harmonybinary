package test

import (
	"fmt"
	"github.com/CarrotVegeta/harmonybinary/hap"
	"testing"
)

func TestHap(t *testing.T) {
	a, err := hap.OpenFile("./c.hap")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
