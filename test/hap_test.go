package test

import (
	"fmt"
	"harmonybinary/app"
	"testing"
)

func TestHap(t *testing.T) {
	a, err := app.OpenFile("./c.hap")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
