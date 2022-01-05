package test

import (
	"fmt"
	"harmonybinary/app"
	"testing"
)

func TestApp(t *testing.T) {
	a, err := app.OpenFile("./123.app")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
