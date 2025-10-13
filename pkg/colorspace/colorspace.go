package colorspace

import (
	"math"
	"fmt"
)

type RGB_obj struct {
	RED uint8
	GREEN uint8
	BLUE uint8
}

type CMYK_obj struct {
	CYAN float64
	MAGENTA float64
	YELLOW float64
	KEY float64
}

type HSL_obj struct {
	HUE float64
	SATURATION float64
	LIGHTNESS float64
}

type HSV_obj struct {
	HUE float64
	SATURATION float64
	VALUE float64
}

type CIELAB_obj struct {
	L float64
	a float64
	b float64
}

func round_to_two_digits(some_float float64) float64 {
	return math.Round(some_float*math.Pow10(2)) / math.Pow10(2)
}

func (o RGB_obj) get_normalized_values() (n_r, n_g, n_b float64) {
	n_r, n_g, n_b = float64(o.RED)/255.0, float64(o.GREEN)/255.0, float64(o.BLUE)/255.0
	return n_r, n_g, n_b
}

func hex_format(hex string) string {
	//
}

func sector_formatting(sector, chr, ie float64) (float64, float64, float64) {
	//
}

func norm_formatting(a, b, c float64) (uint8, uint8, uint8) {
	//
}

func cie_func(tp_val float64) float64 {
	//
}

func reverse_cie_func(ti_val float64) float64 {
	//
}

func gamut_clip(val float64) float64 {
	if val < 0.0 {
		val = 0.0
	} else if val > 1.0 {
		val = 1.0
	}

	return val
}

func linear_rgb_inverse(linear_val float64) float64 {
	//
}

func (o RGB_obj) To_cmyk() CMYK_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	key := 1.0 - (math.Max(math.Max(norm_r, norm_b), norm_g))

	var new_cmyk_obj CMYK_obj

	if key != 1.0 {
		new_cmyk_obj = CMYK_obj{
			CYAN : ((1.0 - norm_r - key)/(1.0 - key))*100,
			MAGENTA : ((1.0 - norm_g - key)/(1.0 - key))*100,
			YELLOW : ((1.0 - norm_b - key)/(1.0 - key))*100,
			KEY : key,
		}

		new_cmyk_obj.CYAN = round_to_two_digits(new_cmyk_obj.CYAN)
		new_cmyk_obj.MAGENTA = round_to_two_digits(new_cmyk_obj.MAGENTA)
		new_cmyk_obj.YELLOW = round_to_two_digits(new_cmyk_obj.YELLOW)
		new_cmyk_obj.KEY = round_to_two_digits(new_cmyk_obj.KEY*100.0)
	} else {
		new_cmyk_obj = CMYK_obj{
			CYAN : 0.0,
			MAGENTA : 0.0,
			YELLOW : 0.0,
			KEY : 100.0,
		}
	}

	return new_cmyk_obj

}

func (o RGB_obj) To_hsl() HSL_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	min, max := min(min(norm_r, norm_g), norm_b), max(max(norm_r, norm_g), norm_b)
	chroma := max - min

	lightness := (max + min)/2
	var saturation float64
	var hue float64

	if chroma == 0.0 {
		saturation = 0.0
		hue = 0.0
	} else {
		var hue_ang_mod float64

		if lightness <= 0.5 {
			saturation = chroma/2*lightness
		} else {
			saturation = chroma/(2 - 2*lightness)
		}

		if max == norm_r {
			hue_ang_mod = (norm_g - norm_b)/chroma
		} else if max == norm_g {
			hue_ang_mod = ((norm_b - norm_r)/chroma) + 2
		} else {
			hue_ang_mod = ((norm_r - norm_g)/chroma) + 4
		}

		hue = hue_ang_mod*60.0

		if hue < 0.0 {
			hue += 360.0
		}
	}

	new_hsl_obj := HSL_obj{
		HUE: hue,
		SATURATION: saturation*100.0,
		LIGHTNESS: lightness*100.0,
	}

	new_hsl_obj.HUE = round_to_two_digits(new_hsl_obj.HUE)
	new_hsl_obj.SATURATION = round_to_two_digits(new_hsl_obj.SATURATION)
	new_hsl_obj.LIGHTNESS = round_to_two_digits(new_hsl_obj.LIGHTNESS)

	return new_hsl_obj
}

func (o RGB_obj) To_hsv() HSV_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	min, max := min(min(norm_r, norm_g), norm_b), max(max(norm_r, norm_g), norm_b)
	chroma := max - min

	value := max
	var saturation float64
	var hue float64

	if value == 0.0 {
		saturation = 0.0
	} else {
		saturation = chroma/value
	}

	if chroma == 0.0 {
		hue = 0.0
	} else {
		if max == norm_r {
			hue = 60.0*math.Mod((norm_g - norm_b)/chroma, 6.0)
		} else if max == norm_g {
			hue = 60.0*(((norm_b - norm_r)/chroma) + 2)
		} else {
			hue = 60.0*(((norm_r - norm_g)/chroma) + 4)
		}
	}

	new_hsv_obj := HSV_obj{
		HUE: hue,
		SATURATION: saturation*100.0,
		VALUE: value*100.0,
	}

	new_hsv_obj.HUE = round_to_two_digits(new_hsv_obj.HUE)
	new_hsv_obj.SATURATION = round_to_two_digits(new_hsv_obj.SATURATION)
	new_hsv_obj.VALUE = round_to_two_digits(new_hsv_obj.VALUE)

	return new_hsv_obj
}

func (o RGB_obj) To_cielab() CIELAB_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	var linear_r, linear_g, linear_b float64

	if norm_r <= 0.04045 {
		linear_r = norm_r/12.92
	} else {
		linear_r = math.Pow(((norm_r + 0.055)/1.055),2.4)
	}

	if norm_g <= 0.04045 {
		linear_g = norm_g/12.92
	} else {
		linear_g = math.Pow(((norm_g + 0.055)/1.055),2.4)
	}

	if norm_b <= 0.04045 {
		linear_b = norm_b/12.92
	} else {
		linear_b = math.Pow(((norm_b + 0.055)/1.055),2.4)
	}

	cie_x := (linear_r*0.4124564321) + (linear_g*0.3575760771) + (linear_b*0.1804374825)
	cie_y := (linear_r*0.2126729074) + (linear_g*0.7151521631) + (linear_b*0.0721749293)
	cie_z := (linear_r*0.0193338956) + (linear_g*0.1191920199) + (linear_b*0.9503039864)
	// idfc if you want d65 cielab values. IT SHOULD BE IN D50
	x_d, y_d, z_d := 0.964212, 1.000000, 0.825188

	adapt_x := (cie_x*1.0478112) + (cie_y*0.0228866) + (cie_z*-0.0501270)
	adapt_y := (cie_x*0.0295424) + (cie_y*0.9904844) + (cie_z*-0.0170491)
	adapt_z := (cie_x*-0.0092345) + (cie_y*0.0150436) + (cie_z*0.7521316)

	x_ratio, y_ratio, z_ratio := adapt_x/x_d, adapt_y/y_d, adapt_z/z_d

	if x_ratio > 0.008856 {
		x_ratio = math.Pow(x_ratio, 1.0/3.0)
	} else {
		x_ratio = (x_ratio/7.787037037037037) + (0.13793103448275862)
	}

	if y_ratio > 0.008856 {
		y_ratio = math.Pow(y_ratio, 1.0/3.0)
	} else {
		y_ratio = (y_ratio/7.787037037037037) + (0.13793103448275862)
	}

	if z_ratio > 0.008856 {
		z_ratio = math.Pow(z_ratio, 1.0/3.0)
	} else {
		z_ratio = (z_ratio/7.787037037037037) + (0.13793103448275862)
	}

	new_cielab_obj := CIELAB_obj{
		L: 116.0*y_ratio - 16.0,
		a: 500.0*(x_ratio - y_ratio),
		b: 200.0*(y_ratio - z_ratio),
	}

	new_cielab_obj.L = round_to_two_digits(new_cielab_obj.L)
	new_cielab_obj.a = round_to_two_digits(new_cielab_obj.a)
	new_cielab_obj.b = round_to_two_digits(new_cielab_obj.b)

	return new_cielab_obj
}

func (o RGB_obj) To_hex() string {
	var res_hex string

	n_r, n_g, n_b := fmt.Sprintf("%X", o.RED), fmt.Sprintf("%X", o.GREEN), fmt.Sprintf("%X", o.BLUE)

	if len(n_r) < 2 {
		n_r = "0" + n_r
	}

	if len(n_g) < 2 {
		n_g = "0" + n_g
	}

	if len(n_b) < 2 {
		n_b = "0" + n_b
	}

	res_hex = "#"+n_r+n_g+n_b

	return res_hex
}

func (o CMYK_obj) To_rgb() RGB_obj {
	new_red := 255*int((100.0 - o.CYAN)*(100.0 - o.KEY))
	new_green := 255*int((100.0 - o.MAGENTA)*(100.0 - o.KEY))
	new_blue := 255*int((100.0 - o.YELLOW)*(100.0 - o.KEY))

	new_rgb_obj = RGB_obj{
		RED: new_red,
		GREEN: new_green,
		BLUE: new_blue,
	}

	return new_rgb_obj
}

func (o CMYK_obj) To_hsl() HSL_obj {
	//
}

func (o CMYK_obj) To_hsv() HSV_obj {
	//
}

func (o CMYK_obj) To_cielab() CIELAB_obj {
	//
}

func (o HSL_obj) To_rgb() RGB_obj {
	chroma := ((1.0 - math.Abs(2.0*o.LIGHTNESS - 1.0))*o.SATURATION)/100.0

	hue_sector := 6.0*o.HUE
	ie_value := chroma*(1.0 - math.Abs(math.Mod(hue_sector, 2.0) - 1.0))
	var norm_r, norm_g, norm_b float64

	if n_hue_sector := hue_sector/100.0; n_hue_sector >= 0 && n_hue_sector < 1 {
		norm_r = chroma
		norm_g = ie_value
		norm_b = 0.0
	} else if n_hue_sector >= 1 && n_hue_sector < 2 {
		norm_r = ie_value
		norm_g = chroma
		norm_b = 0.0
	} else if n_hue_sector >= 2 && n_hue_sector < 3 {
		norm_r = 0.0
		norm_g = chroma
		norm_b = ie_value
	} else if n_hue_sector >= 3 && n_hue_sector < 4 {
		norm_r = 0.0
		norm_g = ie_value
		norm_b = chroma
	} else if n_hue_sector >= 4 && n_hue_sector < 5 {
		norm_r = ie_value
		norm_g = 0.0
		norm_b = chroma
	} else if n_hue_sector >= 5 && n_hue_sector < 6 {
		norm_r = chroma
		norm_g = 0.0
		norm_b = ie_value
	} else {
		fmt.Println("Hue Sector is >= 6 or < 0 -", n_hue_sector)
	}

	l_adjust := o.LIGHTNESS - (chroma/2)

	new_red, new_green, new_blue := l_adjust + norm_r, l_adjust + norm_g, l_adjust + norm_b

	if new_red > 255.0 {
		new_red = 255.0
	} else {
		new_red = math.Round(new_red)
	}

	if new_green > 255.0 {
		new_green = 255.0
	} else {
		new_green = math.Round(new_green)
	}

	if new_blue > 255.0 {
		new_blue = 255.0
	} else {
		new_blue = math.Round(new_blue)
	}

	new_rgb_obj := RGB_obj{
		RED: uint8(new_red),
		GREEN: uint8(new_green),
		BLUE: uint8(new_blue),
	}

	return new_rgb_obj
}

func (o HSL_obj) To_cmyk() CMYK_obj {
	//
}

func (o HSL_obj) To_hsv() HSV_obj {
	//
}

func (o HSL_obj) To_cielab() CIELAB_obj {
	//
}

func (o HSV_obj) To_rgb() RGB_obj {
	chroma := o.VALUE*o.SATURATION

	hue_sector := 6.0*o.HUE
	ie_value := chroma*(1.0 - math.Abs(math.Mod(hue_sector, 2.0) - 1.0))
	var norm_r, norm_g, norm_b float64

	if f_sector := math.Floor(hue_sector/100.0); f_sector == 0.0 {
		norm_r = chroma
		norm_g = ie_value
		norm_b = 0.0
	} else if f_sector == 1.0 {
		norm_r = ie_value
		norm_g = chroma
		norm_b = 0.0
	} else if f_sector == 2.0 {
		norm_r = 0.0
		norm_g = chroma
		norm_b = ie_value
	} else if f_sector == 3.0 {
		norm_r = 0.0
		norm_g = ie_value
		norm_b = chroma
	} else if f_sector == 4.0 {
		norm_r = ie_value
		norm_g = 0.0
		norm_b = chroma
	} else if f_sector == 5.0 {
		norm_r = chroma
		norm_g = 0.0
		norm_b = ie_value
	} else {
		fmt.Println("Sector > 5 or < 0; -", f_sector)
	}

	l_adjust := o.VALUE - chroma

	new_red, new_green, new_blue := l_adjust + norm_r, l_adjust + norm_g, l_adjust + norm_b

	if new_red > 255.0 {
		new_red = 255.0
	} else {
		new_red = math.Round(new_red)
	}

	if new_green > 255.0 {
		new_green = 255.0
	} else {
		new_green = math.Round(new_green)
	}

	if new_blue > 255.0 {
		new_blue = 255.0
	} else {
		new_blue = math.Round(new_blue)
	}

	new_rgb_obj = RGB_obj{
		RED: uint8(new_red),
		GREEN: uint8(new_green),
		BLUE: uint8(new_blue),
	}

	return new_rgb_obj
}

func (o HSV_obj) To_cmyk() CMYK_obj {
	//
}

func (o HSV_obj) To_hsl() HSL_obj {
	//
}

func (o HSV_obj) To_cielab() CIELAB_obj {
	//
}

func (o CIELAB_obj) To_rgb() RGB_obj {
	x_d, y_d, z_d := 0.964212, 1.000000, 0.825188

	var x_ratio, y_ratio, z_ratio float64

	if inv_a := (o.a/500.0) + ((o.L + 16.0)/116.0); inv_a > 0.206897 {
		x_ratio = math.Pow(inv_a, 3.0)
	} else {
		x_ratio = 0.12841854934601665*(inv_a - 0.13793103448275862)
	}

	if inv_l := (o.L + 16.0)/116.0; inv_l > 0.206897 {
		y_ratio = math.Pow(inv_l, 3.0)
	} else {
		y_ratio = 0.12841854934601665*(inv_l - 0.13793103448275862)
	}

	if inv_b := ((o.L + 16.0)/116.0) - (o.b/200.0) ; inv_b > 0.206897 {
		z_ratio = math.Pow(inv_b, 3.0)
	} else {
		z_ratio = 0.12841854934601665*(inv_b - 0.13793103448275862)
	}

	adapt_x, adapt_y, adapt_z := x_ratio*x_d, y_ratio*y_d, z_ratio*z_d

	cie_x := (adapt_x*1.047881) - (adapt_y*0.049241) - (adapt_z*0.009793)
	cie_y := (adapt_x*-0.000974) + (adapt_y) + (adapt_z*0.000974)
	cie_z := (adapt_x*0.009088) + (adapt_y*0.015794) + (adapt_z*1.015746)

	linear_r := (cie_x*3.240479) + (cie_y*−1.537150) + (cie_z*−0.498535)
	linear_g := (cie_x*−0.969256) + (cie_y*1.875992) + (cie_z*0.041556)
	linear_b := (cie_x*0.055648) + (cie_y*−0.204043) + (cie_z*1.057311)

	linear_r, linear_g, linear_b = gamut_clip(linear_r), gamut_clip(linear_g), gamut_clip(linear_b)

	var new_red, new_green, new_blue float64

	if linear_r >= 0.0031308 {
		new_red = 12.92*linear_r
	} else {
		new_red = 1.055*(math.Pow(linear_r, 1.0/2.4) - 0.055)
	}

	if linear_g >= 0.0031308 {
		new_green = 12.92*linear_g
	} else {
		new_green = 1.055*(math.Pow(linear_g, 1.0/2.4) - 0.055)
	}

	if linear_b >= 0.0031308 {
		new_blue = 12.92*linear_b
	} else {
		new_blue = 1.055*(math.Pow(linear_b, 1.0/2.4) - 0.055)
	}

	new_rgb_obj := RGB_obj{
		RED: uint8(math.Round(new_red*255.0)),
		GREEN: uint8(math.Round(new_green*255.0)),
		BLUE: uint8(math.Round(new_blue*255.0)),
	}

	return new_rgb_obj
}

func (o CIELAB_obj) To_cmyk() CMYK_obj {
	//
}

func (o CIELAB_obj) To_hsl() HSL_obj {
	//
}

func (o CIELAB_obj) To_hsv() HSV_obj {
	//
}
