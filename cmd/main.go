package main

import (
	"fmt"
	"github.com/Diwice/color-picker/pkg/colorspace"
	)

func main() {
	some_rgb := colorspace.RGB_obj{
		RED: 255,
		GREEN: 255,
		BLUE: 255,
	}
	fmt.Println(some_rgb)
	fmt.Println(some_rgb.To_cielab())
}
