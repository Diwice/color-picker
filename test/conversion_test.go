package colorspace

import (
	"testing"
	"github.com/Diwice/color-picker/pkg/colorspace"
)

func Test_RGB_to_Hex(t *testing.T) {
	rgb_obj := colorspace.RGB_obj{
		RED: 1,
		GREEN: 2,
		BLUE: 3,
	}

	converted_obj := rgb_obj.To_hex()

	if converted_obj != "#010203" {
		t.Errorf("Expected : #010203 ; Got : %s", converted_obj)
	}
}

func Test_RGB_to_CMYK(t *testing.T) {
	rgb_obj := colorspace.RGB_obj{
		RED: 3,
		GREEN: 8,
		BLUE: 163,
	}

	cmyk_obj := rgb_obj.To_cmyk()

	if cmyk_obj.CYAN != 98.16 || cmyk_obj.MAGENTA != 95.09 || cmyk_obj.YELLOW != 0.0 || cmyk_obj.KEY != 36.08 {
		t.Errorf("Expected : C-98.16/M-95.09/Y-0.0/K-36.08 ; Got : C-%f/M-%f/Y-%f/K-%f", cmyk_obj.CYAN, cmyk_obj.MAGENTA, cmyk_obj.YELLOW, cmyk_obj.KEY)
	}
}

func Test_RGB_to_HSV(t *testing.T) {
	rgb_obj := colorspace.RGB_obj{
		RED: 25,
		GREEN: 15,
		BLUE: 10,
	}

	hsv_obj := rgb_obj.To_hsv()

	if hsv_obj.HUE != 20.0 || hsv_obj.SATURATION != 60.0 || hsv_obj.VALUE != 9.8 {
		t.Errorf("Expected : H-20.0/S-60.0/V-9.8 ; Got : H-%f/S-%f/V-%f", hsv_obj.HUE, hsv_obj.SATURATION, hsv_obj.VALUE)
	}
}

func Test_RGB_to_HSL(t *testing.T) {
	rgb_obj := colorspace.RGB_obj{
		RED: 50,
		GREEN: 100,
		BLUE: 150,
	}

	hsl_obj := rgb_obj.To_hsl()

	if hsl_obj.HUE != 210.0 || hsl_obj.SATURATION != 50.0 || hsl_obj.LIGHTNESS != 39.22 {
		t.Errorf("Expected : H-210.0/S-50.0/L-39.22 ; Got : H-%f/S-%f/L-%f", hsl_obj.HUE, hsl_obj.SATURATION, hsl_obj.LIGHTNESS)
	}
}

func Test_RGB_to_CIELAB(t *testing.T) {
	rgb_obj := colorspace.RGB_obj{
		RED: 255,
		GREEN: 255,
		BLUE: 255,
	}

	cielab_obj := rgb_obj.To_cielab()

	if cielab_obj.L != 100.0 || cielab_obj.A != 0.0 || cielab_obj.B != 0.0 {
		t.Errorf("Expected : L-100.0/a-0.0/b-0.0 ; Got : L-%f/a-%f/b-%f", cielab_obj.L, cielab_obj.A, cielab_obj.B)
	}
}

/* Since RGB conversions are fully tested there's no point in testing 
most cases but conversion to RGB (HSL to HSV and vice versa tested separately)

func Test_CMYK_to_RGB(t *testing.T) {
	//
}

func Test_CMYK_RGB_Gamut_Clip(t *testing.T) {
	//
}

func Test_HSV_to_RGB(t *testing.T) {
	//
}

func Test_HSV_to_HSL(t *testing.T) {
	//
}

func Test_HSL_to_RGB(t *testing.T) {
	//
}

func Test_HSL_to_HSV(t *testing.T) {
	//
}

func Test_CIELAB_to_RGB(t *testing.T) {
	//
}

func Test_CIELAB_RGB_Gamut_Clip(t *testing.T){
	//
} */
