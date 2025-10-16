package colorspace

import (
	"math"
	"fmt"
)

/* Since I'm stumbling on padding anyways, decided to use float64 data types instead.
They produce more precise values, which is crucial when working with colors... */

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

// Helper functions. Made for "DRY".

func round_to_two_digits(some_float float64) float64 {
	return math.Round(some_float*math.Pow10(2)) / math.Pow10(2)
}

func hex_format(hex string) string {
	if len(hex) < 2 {
		return "0" + hex
	}

	return hex
}

// Both HSL/HSV to RGB are using the same algorithm

func sector_formatting(sector, chr, ie float64) (float64, float64, float64, error) {
	var n_r, n_g, n_b float64

	if n_hue_sector := sector/100.0; n_hue_sector >= 0 && n_hue_sector < 1 {
		n_r = chr
		n_g = ie
		n_b = 0.0
	} else if n_hue_sector >= 1 && n_hue_sector < 2 {
		n_r = ie
		n_g = chr
		n_b = 0.0
	} else if n_hue_sector >= 2 && n_hue_sector < 3 {
		n_r = 0.0
		n_g = chr
		n_b = ie
	} else if n_hue_sector >= 3 && n_hue_sector < 4 {
		n_r = 0.0
		n_g = ie
		n_b = chr
	} else if n_hue_sector >= 4 && n_hue_sector < 5 {
		n_r = ie
		n_g = 0.0
		n_b = chr
	} else if n_hue_sector >= 5 && n_hue_sector < 6 {
		n_r = chr
		n_g = 0.0
		n_b = ie
	} else {
		return 0.0, 0.0, 0.0, fmt.Errorf("Sector out of range : %f", n_hue_sector)
	}

	return n_r, n_g, n_b, nil
}

/* new_formatting INCLUDES gamut clipping.
Both HSV and HSL spaces hold more color information than RGB space does.
Causes data color loss. Not like I could directly convert some color spaces to others. */

func new_formatting(a, b, c float64) (uint8, uint8, uint8) {
	if a > 255.0 {
		a = 255.0
	} else {
		a = math.Round(a)
	}

	if b > 255.0 {
		b = 255.0
	} else {
		b = math.Round(b)
	}

	if c > 255.0 {
		c = 255.0
	} else {
		c = math.Round(c)
	}

	return uint8(a), uint8(b), uint8(c)
}

func cie_func(tp_val float64) float64 {
	var sub_val float64

	if tp_val > 0.008856 {
		sub_val = math.Pow(tp_val, 1.0/3.0)
	} else {
		sub_val = (tp_val/7.787037037037037) + (0.13793103448275862)
	}

	return sub_val
}

func reverse_cie_func(ti_val float64) float64 {
	var sub_val float64

	if ti_val > 0.206897 {
		sub_val = math.Pow(ti_val, 3.0)
	} else {
		sub_val = 0.12841854934601665*(ti_val - 0.13793103448275862)
	}

	return sub_val
}

func gamut_clip(val float64) float64 {
	if val < 0.0 {
		val = 0.0
	} else if val > 1.0 {
		val = 1.0
	}

	return val
}

func norm_to_linear(norm_val float64) float64 {
	var linear_color float64

	if norm_val <= 0.04045 {
		linear_color = norm_val/12.92
	} else {
		linear_color = math.Pow(((norm_val + 0.055)/1.055), 2.4)
	}

	return linear_color
}

func linear_rgb_inverse(linear_val float64) float64 {
	var inversed_color float64

	if linear_val >= 0.0031308 {
		inversed_color = 12.92*linear_val
	} else {
		inversed_color = 1.055*(math.Pow(linear_val, 1.0/2.4) - 0.055)
	}

	return inversed_color
}

func (o RGB_obj) get_normalized_values() (n_r, n_g, n_b float64) {
	n_r, n_g, n_b = float64(o.RED)/255.0, float64(o.GREEN)/255.0, float64(o.BLUE)/255.0
	return n_r, n_g, n_b
}

// Not gonna explain anything but RGB to CIE Lab since those are basic conversions.

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

/* Rough outline on RGB to CIE Lab algorithm :
RGB -> normalized RGB -> linear RGB -> CIE XYZ D65 -> CIE XYZ D50 ->
White Point to XYZ ratio -> CIE function conversions on ratios -> CIE Lab */

func (o RGB_obj) To_cielab() CIELAB_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()

	linear_r, linear_g, linear_b := norm_to_linear(norm_r), norm_to_linear(norm_g), norm_to_linear(norm_b)
	/* Numbers which are used in those multiplications come from mathematical definition of RGB color space.
	They are a part of matrix with coefficients used for RGB -> CIE XYZ conversions */
	cie_x := (linear_r*0.4124564321) + (linear_g*0.3575760771) + (linear_b*0.1804374825)
	cie_y := (linear_r*0.2126729074) + (linear_g*0.7151521631) + (linear_b*0.0721749293)
	cie_z := (linear_r*0.0193338956) + (linear_g*0.1191920199) + (linear_b*0.9503039864)

	/* Further calculations mutate XYZ D65 values to XYZ D50 values, since RGB is defined relatively to D65.
	If you want D65 CIE Lab values - remove everything up to adapt values and replace d-values.
	Then replace adapt values with cie values in ratio values.*/

	// White Point D50 values
	x_d, y_d, z_d := 0.964212, 1.000000, 0.825188
	// Adapting XYZ D65 to D50. Bradford CAT matrix values.
	adapt_x := (cie_x*1.0478112) + (cie_y*0.0228866) + (cie_z*-0.0501270)
	adapt_y := (cie_x*0.0295424) + (cie_y*0.9904844) + (cie_z*-0.0170491)
	adapt_z := (cie_x*-0.0092345) + (cie_y*0.0150436) + (cie_z*0.7521316)

	x_ratio, y_ratio, z_ratio := adapt_x/x_d, adapt_y/y_d, adapt_z/z_d

	final_x, final_y, final_z := cie_func(x_ratio), cie_func(y_ratio), cie_func(z_ratio)

	new_cielab_obj := CIELAB_obj{
		L: 116.0*final_y - 16.0,
		a: 500.0*(final_x - final_y),
		b: 200.0*(final_y - final_z),
	}

	new_cielab_obj.L = round_to_two_digits(new_cielab_obj.L)
	new_cielab_obj.a = round_to_two_digits(new_cielab_obj.a)
	new_cielab_obj.b = round_to_two_digits(new_cielab_obj.b)

	return new_cielab_obj
}

func (o RGB_obj) To_hex() string {
	var res_hex string

	hex_r, hex_g, hex_b := fmt.Sprintf("%X", o.RED), fmt.Sprintf("%X", o.GREEN), fmt.Sprintf("%X", o.BLUE)

	hex_r, hex_g, hex_b = hex_format(hex_r), hex_format(hex_g), hex_format(hex_b)

	res_hex = "#"+hex_r+hex_g+hex_b

	return res_hex
}

func (o CMYK_obj) To_rgb() RGB_obj {
	new_red := uint8(255.0*(((100.0 - o.CYAN)*(100.0 - o.KEY))/100.0))
	new_green := uint8(255.0*(((100.0 - o.MAGENTA)*(100.0 - o.KEY))/100.0))
	new_blue := uint8(255.0*(((100.0 - o.YELLOW)*(100.0 - o.KEY))/100.0))

	new_rgb_obj := RGB_obj{
		RED: new_red,
		GREEN: new_green,
		BLUE: new_blue,
	}

	return new_rgb_obj
}

/* Couldn't find direct conversion formulas myself.
For some of the further transformations RGB is used as an intermidiate value.
Color data is getting lost in process. */

func (o CMYK_obj) To_hsl() HSL_obj {
	new_rgb_obj := o.To_rgb()

	new_hsl_obj := new_rgb_obj.To_hsl()

	return new_hsl_obj
}

func (o CMYK_obj) To_hsv() HSV_obj {
	new_rgb_obj := o.To_rgb()

	new_hsv_obj := new_rgb_obj.To_hsv()

	return new_hsv_obj
}

func (o CMYK_obj) To_cielab() CIELAB_obj {
	new_rgb_obj := o.To_rgb()

	new_cielab_obj := new_rgb_obj.To_cielab()

	return new_cielab_obj
}

func (o HSL_obj) To_rgb() RGB_obj {
	chroma := ((1.0 - math.Abs(2.0*(o.LIGHTNESS/100.0) - 1.0))*(o.SATURATION/100.0))/100.0

	hue_sector := 6.0*o.HUE
	ie_value := chroma*(1.0 - math.Abs(math.Mod(hue_sector, 2.0) - 1.0))
	var norm_r, norm_g, norm_b float64

	if sub_r, sub_g, sub_b, err := sector_formatting(hue_sector, chroma, ie_value); err == nil {
		norm_r = sub_r
		norm_g = sub_g
		norm_b = sub_b
	} else {
		fmt.Printf("Error converting HSL to RGB : %ss / %s\n", err)
		return RGB_obj{RED: 0.0, GREEN: 0.0, BLUE: 0.0}
	}

	l_adjust := (o.LIGHTNESS/100.0) - (chroma/2.0)

	sub_red, sub_green, sub_blue := l_adjust + norm_r, l_adjust + norm_g, l_adjust + norm_b

	new_red, new_green, new_blue := new_formatting(sub_red, sub_green, sub_blue)

	new_rgb_obj := RGB_obj{
		RED: new_red,
		GREEN: new_green,
		BLUE: new_blue,
	}

	return new_rgb_obj
}

func (o HSL_obj) To_cmyk() CMYK_obj {
	new_rgb_obj := o.To_rgb()

	new_cmyk_obj := new_rgb_obj.To_cmyk()

	return new_cmyk_obj
}

func (o HSL_obj) To_hsv() HSV_obj {
	value := (o.LIGHTNESS/100.0) + (o.SATURATION/100.0)*min((o.LIGHTNESS/100.0), 1 - (o.LIGHTNESS/100.0))

	var saturation float64

	if value > 0.0 {
		saturation = 2*(value - (o.LIGHTNESS/100.0))/value
	} else if value == 0.0 {
		saturation = 0.0
	}

	new_hsv_obj := HSV_obj{
		HUE: o.HUE,
		SATURATION: saturation*100.0,
		VALUE: value*100.0,
	}

	new_hsv_obj.SATURATION = round_to_two_digits(new_hsv_obj.SATURATION)
	new_hsv_obj.VALUE = round_to_two_digits(new_hsv_obj.VALUE)

	return new_hsv_obj
}

func (o HSL_obj) To_cielab() CIELAB_obj {
	new_rgb_obj := o.To_rgb()

	new_cielab_obj := new_rgb_obj.To_cielab()

	return new_cielab_obj
}

func (o HSV_obj) To_rgb() RGB_obj {
	chroma := (o.VALUE/100.0)*(o.SATURATION/100.0)

	hue_sector := 6.0*o.HUE
	ie_value := chroma*(1.0 - math.Abs(math.Mod(hue_sector, 2.0) - 1.0))
	var norm_r, norm_g, norm_b float64

	if sub_r, sub_g, sub_b, err := sector_formatting(hue_sector, chroma, ie_value); err == nil {
		norm_r = sub_r
		norm_g = sub_g
		norm_b = sub_b
	} else {
		fmt.Printf("Error converting HSL to RGB : %ss / %s\n", err)
		return RGB_obj{RED: 0.0, GREEN: 0.0, BLUE: 0.0}
	}

	l_adjust := (o.VALUE/100.0) - chroma

	sub_red, sub_green, sub_blue := l_adjust + norm_r, l_adjust + norm_g, l_adjust + norm_b

	new_red, new_green, new_blue := new_formatting(sub_red, sub_green, sub_blue)

	new_rgb_obj := RGB_obj{
		RED: new_red,
		GREEN: new_green,
		BLUE: new_blue,
	}

	return new_rgb_obj
}

func (o HSV_obj) To_cmyk() CMYK_obj {
	new_rgb_obj := o.To_rgb()

	new_cmyk_obj := new_rgb_obj.To_cmyk()

	return new_cmyk_obj
}

func (o HSV_obj) To_hsl() HSL_obj {
	lightness := (o.VALUE/100.0)*(1 - ((o.SATURATION/100.0)/2.0))

	var saturation float64

	if lightness > 0 && lightness < 1 {
		saturation = ((o.VALUE/100.0) - lightness)/min(lightness, 1.0 - lightness)
	} else if lightness == 0 || lightness == 1 {
		saturation = 0.0
	}

	new_hsl_obj := HSL_obj{
		HUE: o.HUE,
		SATURATION: saturation*100.0,
		LIGHTNESS: lightness*100.0,
	}

	new_hsl_obj.SATURATION = round_to_two_digits(new_hsl_obj.SATURATION)
	new_hsl_obj.LIGHTNESS = round_to_two_digits(new_hsl_obj.LIGHTNESS)

	return new_hsl_obj
}

func (o HSV_obj) To_cielab() CIELAB_obj {
	new_rgb_obj := o.To_rgb()

	new_cielab_obj := new_rgb_obj.To_cielab()

	return new_cielab_obj
}

// CIE Lab (D50) to RGB algorithm is pretty much the same but inversed RGB to CIE Lab (D50) 

func (o CIELAB_obj) To_rgb() RGB_obj {
	x_d, y_d, z_d := 0.964212, 1.000000, 0.825188

	inv_a := (o.a/500.0) + ((o.L + 16.0)/116.0)
	x_ratio := reverse_cie_func(inv_a)

	inv_l := (o.L + 16.0)/116.0
	y_ratio := reverse_cie_func(inv_l)

	inv_b := ((o.L + 16.0)/116.0) - (o.b/200.0)
	z_ratio := reverse_cie_func(inv_b)

	adapt_x, adapt_y, adapt_z := x_ratio*x_d, y_ratio*y_d, z_ratio*z_d

	cie_x := (adapt_x*1.047881) - (adapt_y*0.049241) - (adapt_z*0.009793)
	cie_y := (adapt_x*-0.000974) + (adapt_y) + (adapt_z*0.000974)
	cie_z := (adapt_x*0.009088) + (adapt_y*0.015794) + (adapt_z*1.015746)

	linear_r := (cie_x*3.240479) + (cie_y*-1.537150) + (cie_z*-0.498535)
	linear_g := (cie_x*-0.969256) + (cie_y*1.875992) + (cie_z*0.041556)
	linear_b := (cie_x*0.055648) + (cie_y*-0.204043) + (cie_z*1.057311)

	linear_r, linear_g, linear_b = gamut_clip(linear_r), gamut_clip(linear_g), gamut_clip(linear_b)

	new_red, new_green, new_blue := linear_rgb_inverse(linear_r), linear_rgb_inverse(linear_g), linear_rgb_inverse(linear_b)

	new_rgb_obj := RGB_obj{
		RED: uint8(math.Round(new_red*255.0)),
		GREEN: uint8(math.Round(new_green*255.0)),
		BLUE: uint8(math.Round(new_blue*255.0)),
	}

	return new_rgb_obj
}

func (o CIELAB_obj) To_cmyk() CMYK_obj {
	new_rgb_obj := o.To_rgb()

	new_cmyk_obj := new_rgb_obj.To_cmyk()

	return new_cmyk_obj
}

func (o CIELAB_obj) To_hsl() HSL_obj {
	new_rgb_obj := o.To_rgb()

	new_hsl_obj := new_rgb_obj.To_hsl()

	return new_hsl_obj
}

func (o CIELAB_obj) To_hsv() HSV_obj {
	new_rgb_obj := o.To_rgb()

	new_hsv_obj := new_rgb_obj.To_hsv()

	return new_hsv_obj
}
