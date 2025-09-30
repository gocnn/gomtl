package main

import (
	"fmt"

	"github.com/gocnn/gomtl"
)

func main() {
	device := gomtl.CreateSystemDefaultDevice()
	fmt.Println(device)
}
