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
	fmt.Println(some_rgb)
	fmt.Println(some_rgb.To_cmyk())
	fmt.Println(some_rgb.To_hsv())
	fmt.Println(some_rgb.To_hsl())
}
