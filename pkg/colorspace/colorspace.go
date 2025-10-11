package colorspace

import "math"

type RGB_obj struct {
	RED uint8
	BLUE uint8
	GREEN uint8
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
	a int16
	b int16
	L uint8
}

func (o RGB_obj) get_normalized_values() (float64, float64, float64) {
	norm_r, norm_g, norm_b := float64(o.RED)/255.0, float64(o.GREEN)/255.0, float64(o.BLUE)/255.0 //normalized values used across rgb conversion calculated here
	return norm_r, norm_g, norm_b
}

func (o RGB_obj) to_cmyk() CMYK_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	key := 1.0 - (math.Max(math.Max(norm_r, norm_b), norm_g)) // key (black) value

	var new_cmyk_obj CMYK_obj

	if key != 1.0 { // if key is not fully black - other colors are present
		new_cmyk_obj = CMYK_obj{
			CYAN : ((1.0 - norm_r - key)/(1.0 - key))*100,
			MAGENTA : ((1.0 - norm_g - key)/(1.0 - key))*100,
			YELLOW : ((1.0 - norm_b - key)/(1.0 - key))*100,
			KEY : key,
		}
		// round with 2 digit precision for C, M, Y
		new_cmyk_obj.CYAN = math.Round(new_cmyk_obj.CYAN*(math.Pow10(2))) / math.Pow10(2)
		new_cmyk_obj.MAGENTA = math.Round(new_cmyk_obj.MAGENTA*(math.Pow10(2))) / math.Pow10(2)
		new_cmyk_obj.YELLOW = math.Round(new_cmyk_obj.YELLOW*(math.Pow10(2))) / math.Pow10(2)
	} else {
		new_cmyk_obj = CMYK_obj{
			CYAN : 0.0,
			MAGENTA : 0.0,
			YELLOW : 0.0,
			KEY : key,
		}
	}

	return new_cmyk_obj

}

func (o RGB_obj) to_hsl() HSL_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	min, max := min(min(norm_r, norm_g), norm_b), max(max(norm_r, norm_g), norm_b)
	chroma := max - min

	lightness := (max + min)/2 // average of max and min present colors
	var saturation float64
	var hue float64

	if chroma == 0.0 { // if chroma is 0.0 then color is grayscale, meaning saturation and hue are both 0
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
		// if the angle is negative - got to make it positive
		if hue < 0.0 {
			hue += 360.0
		}
	}

	new_hsl_obj := HSL_obj{
		HUE: hue,
		SATURATION: saturation*100.0,
		LIGHTNESS: lightness*100.0,
	}
	// won't mention it again - round with 2 digit precision
	new_hsl_obj.HUE = math.Round(new_hsl_obj.HUE*(math.Pow10(2))) / math.Pow10(2)
	new_hsl_obj.SATURATION = math.Round(new_hsl_obj.SATURATION*(math.Pow10(2))) / math.Pow10(2)
	new_hsl_obj.LIGHTNESS = math.Round(new_hsl_obj.LIGHTNESS*(math.Pow10(2))) / math.Pow10(2)

	return new_hsl_obj
}

func (o RGB_obj) to_hsv() HSV_obj {
	norm_r, norm_g, norm_b := o.get_normalized_values()
	min, max := min(min(norm_r, norm_g), norm_b), max(max(norm_r, norm_g), norm_b)
	chroma := max - min

	value := max // value is equals to brightness, which is max component
	var saturation float64
	var hue float64

	if value == 0.0 {
		saturation = 0.0
	} else {
		saturation = chroma/value // saturation is a ratio of chroma to value
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

	new_hsv_obj.HUE = math.Round(new_hsv_obj.HUE*(math.Pow10(2))) / math.Pow10(2)
	new_hsv_obj.SATURATION = math.Round(new_hsv_obj.SATURATION*(math.Pow10(2))) / math.Pow10(2)
	new_hsv_obj.VALUE = math.Round(new_hsv_obj.VALUE*(math.Pow10(2))) / math.Pow10(2)

	return new_hsv_obj
}
