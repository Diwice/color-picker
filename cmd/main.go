package main

import (
	"fmt"
	"github.com/Diwice/color-picker/pkg/colorspace"
	)

func main() {
	some_rgb := colorspace.RGB_obj{
		RED: 3,
		GREEN: 8,
		BLUE: 163,
	}
	fmt.Println(some_rgb)
	fmt.Println(some_rgb.To_cmyk())
	fmt.Println(some_rgb.To_hsv())
	fmt.Println(some_rgb.To_hsl())
	fmt.Println(some_rgb.To_cielab())
	fmt.Println(some_rgb.To_hex())
}
