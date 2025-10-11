package main

import (
	"fmt"
	"github.com/Diwice/color-picker/pkg/colorspace"
	)

func main() {
	some_rgb := colorspace.RGB_obj{
		RED: 200,
		GREEN: 200,
		BLUE: 100,
	}
	fmt.Print(some_rgb)
	fmt.Print("\n")
}
