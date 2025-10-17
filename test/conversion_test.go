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
		t.Errorf("Expected : C-98.16/M-95.09/Y-0.0/K-36.08 ; Got : C-%.2f/M-%.2f/Y-%.2f/K-%.2f", cmyk_obj.CYAN, cmyk_obj.MAGENTA, cmyk_obj.YELLOW, cmyk_obj.KEY)
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
		t.Errorf("Expected : H-20.0/S-60.0/V-9.8 ; Got : H-%.2f/S-%.2f/V-%.2f", hsv_obj.HUE, hsv_obj.SATURATION, hsv_obj.VALUE)
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
		t.Errorf("Expected : H-210.0/S-50.0/L-39.22 ; Got : H-%.2f/S-%.2f/L-%.2f", hsl_obj.HUE, hsl_obj.SATURATION, hsl_obj.LIGHTNESS)
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
		t.Errorf("Expected : L-100.0/a-0.0/b-0.0 ; Got : L-%.2f/a-%.2f/b-%.2f", cielab_obj.L, cielab_obj.A, cielab_obj.B)
	}
}

/* Since RGB conversions are fully tested there's no point in testing 
most cases but conversion to RGB (HSL to HSV and vice versa tested separately) */

func Test_CMYK_to_RGB(t *testing.T) {
	cmyk_obj := colorspace.CMYK_obj{
		CYAN: 40.0,
		MAGENTA: 30.0,
		YELLOW: 30.0,
		KEY: 10.0,
	}

	rgb_obj := cmyk_obj.To_rgb()

	if rgb_obj.RED != 138 || rgb_obj.GREEN != 161 || rgb_obj.BLUE != 161 {
		t.Errorf("Expected : R-138/G-161/B-161 ; Got : R-%d/G-%d/B-%d", rgb_obj.RED, rgb_obj.GREEN, rgb_obj.BLUE)
	}
}

func Test_CMYK_RGB_Gamut_Clip(t *testing.T) {
	cmyk_obj := colorspace.CMYK_obj{
		CYAN: 100.0,
		MAGENTA: 0.0,
		YELLOW: 0.0,
		KEY: 0.0,
	}

	rgb_obj := cmyk_obj.To_rgb()

	if rgb_obj.RED != 0 || rgb_obj.GREEN != 255 || rgb_obj.BLUE != 255 {
		t.Errorf("Expected : R-0/G-255/B-255 ; Got : R-%d/G-%d/B-%d", rgb_obj.RED, rgb_obj.GREEN, rgb_obj.BLUE)
	}
}

func Test_HSV_to_RGB(t *testing.T) {
	hsv_obj := colorspace.HSV_obj{
		HUE: 210.0,
		SATURATION: 70.0,
		VALUE: 80.0,
	}

	rgb_obj := hsv_obj.To_rgb()

	if rgb_obj.RED != 61 || rgb_obj.GREEN != 133 || rgb_obj.BLUE != 204 {
		t.Errorf("Expected : R-61/G-133/B-204 ; Got : R-%d/G-%d/B-%d", rgb_obj.RED, rgb_obj.GREEN, rgb_obj.BLUE)
	}
}

func Test_HSV_to_HSL(t *testing.T) {
	hsv_obj := colorspace.HSV_obj{
		HUE: 210.0,
		SATURATION: 70.0,
		VALUE: 80.0,
	}

	hsl_obj := hsv_obj.To_hsl()

	if hsl_obj.HUE != 210.0 || hsl_obj.SATURATION != 58.33 || hsl_obj.LIGHTNESS != 52.0 {
		t.Errorf("Expected : H-210.0/S-52.0/L-52.0 ; Got : H-%.2f/S-%.2f/L-%.2f", hsl_obj.HUE, hsl_obj.SATURATION, hsl_obj.LIGHTNESS)
	}
}

func Test_HSL_to_RGB(t *testing.T) {
	hsl_obj := colorspace.HSL_obj{
		HUE: 150.0,
		SATURATION: 80.0,
		LIGHTNESS: 40.0,
	}

	rgb_obj := hsl_obj.To_rgb()

	if rgb_obj.RED != 20 || rgb_obj.GREEN != 184 || rgb_obj.BLUE != 102 {
		t.Errorf("Expected : R-20/G-184/B-102 ; Got : R-%d/G-%d/B-%d", rgb_obj.RED, rgb_obj.GREEN, rgb_obj.BLUE)
	}
}

func Test_HSL_to_HSV(t *testing.T) {
	hsl_obj := colorspace.HSL_obj{
		HUE: 150.0,
		SATURATION: 80.0,
		LIGHTNESS: 40.0,
	}

	hsv_obj := hsl_obj.To_hsv()

	if hsv_obj.HUE != 150.0 || hsv_obj.SATURATION != 88.89 || hsv_obj.VALUE != 72.0 {
		t.Errorf("Expected : H-150.0/S-78.0/V-72.0 ; Got : H-%.2f/S-%.2f/V-%.2f", hsv_obj.HUE, hsv_obj.SATURATION, hsv_obj.VALUE)
	}
}

func Test_CIELAB_to_RGB(t *testing.T) {
	cielab_obj := colorspace.CIELAB_obj{
		L: 60.0,
		A: 25.0,
		B: 50.0,
	}

	rgb_obj := cielab_obj.To_rgb()

	if rgb_obj.RED != 226 || rgb_obj.GREEN != 147 || rgb_obj.BLUE != 0 {
		t.Errorf("Expected : R-226/G-147/B-0 ; Got : R-%d/G-%d/B-%d", rgb_obj.RED, rgb_obj.GREEN, rgb_obj.BLUE)
	}
}

func Test_CIELAB_RGB_Gamut_Clip(t *testing.T){
	cielab_obj := colorspace.CIELAB_obj{
		L: 90.0,
		A: 128.0,
		B: 128.0,
	}

	rgb_obj := cielab_obj.To_rgb()

	if rgb_obj.RED != 255 || rgb_obj.GREEN != 0 || rgb_obj.BLUE != 0 {
		t.Errorf("Expected : R-255/G-0/B-0 ; Got : R-%d/G-%d/B-%d", rgb_obj.RED, rgb_obj.GREEN, rgb_obj.BLUE)
	}
}
