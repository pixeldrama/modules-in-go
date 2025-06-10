package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("example error")
	fmt.Println(err)
}
