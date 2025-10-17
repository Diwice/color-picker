package main

import (
	"fmt"
	"github.com/Diwice/color-picker/pkg/colorspace"
	)

func main() {
	cmyk := colorspace.CMYK_obj{
		CYAN: ,
		MAGENTA: ,
		YELLOW: ,
		KEY: ,
	}
	rgb_zero := cmyk.To_rgb()
	fmt.Print("---CMYK regular---\n", cmyk, "\n", rgb_zero)
	
	hsv := colorspace.HSV_obj{
		HUE: ,
		SATURATION: ,
		VALUE: ,
	}
	rgb_one := hsv.To_rgb()
	fmt.Print("\n---HSV regular---\n", hsv, "\n", rgb_one)

	hsl := colorspace.HSL_obj{
		HUE: ,
		SATURATION: ,
		LIGHTNESS: ,
	}
	rgb_two := hsl.To_rgb()
	fmt.Print("\n---HSL regular---\n", hsl, "\n", rgb_two)

	cielab := colorspace.CIELAB_obj{
		L: ,
		A: ,
		B: ,
	}
	rgb_three := cielab.To_rgb()
	fmt.Print("\n---CIE Lab regular---\n", cielab, "\n", rgb_three)

	cmyk_gamut := colorspace.CMYK_obj{
		CYAN: ,
		MAGENTA: ,
		YELLOW: ,
		KEY: ,
	}
	rgb_gamut_zero := cmyk_gamut.To_rgb()
	fmt.Print("\n---CMYK with gamut---\n", cmyk_gamut, "\n", rgb_gamut_zero)

	cielab_gamut := colorspace.CIELAB_obj{
		L: 94.0,
		A: 128.0,
		B: 128.0,
	}
	rgb_gamut_one := cielab_gamut.To_rgb()
	fmt.Print("\n---CIE Lab with gamut---\n", cielab_gamut, "\n", rgb_gamut_one)

	hsv_to := colorspace.HSV_obj{
		HUE: ,
		SATURATION: ,
		VALUE: ,
	}
	hsl_from := hsv_to.To_hsl()
	fmt.Print("\n---HSV to HSL---\n", hsv_to, "\n", hsl_from)

	hsl_to := colorspace.HSL_obj{
		HUE: ,
		SATURATION: ,
		VALUE: ,
	}
	hsv_from := hsl_to.To_hsv()
	fmt.Print("\n---HSL to HSV---\n", hsl_to, "\n", hsv_from)
}
