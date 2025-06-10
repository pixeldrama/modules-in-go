package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("example error from workspace")
	fmt.Println(err)
}
