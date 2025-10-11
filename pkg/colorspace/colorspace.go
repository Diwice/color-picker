package colorspace

import "math"

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
	// fuck this shit shitty shit math.Pow
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
	// d50 white point values
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
