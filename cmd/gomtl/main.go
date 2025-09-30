package main

import (
	"fmt"

	"github.com/gocnn/gomtl"
)

func main() {
	device, err := gomtl.CreateSystemDefaultDevice()
	if err != nil {
		panic(err)
	}
	fmt.Println(device)
}
