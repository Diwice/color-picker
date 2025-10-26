package main

import (
	"fmt"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/Diwice/color-picker/pkg/colorspace"
)

type color_container struct {
	rgb *colorspace.RGB_obj
	cmyk *colorspace.CMYK_obj
	hsv *colorspace.HSV_obj
	hsl *colorspace.HSL_obj
	lab *colorspace.CIELAB_obj
	hex string
}

func create_empty_container() *color_container {
	starting_color := colorspace.RGB_obj{
		RED: 0,
		GREEN: 0,
		BLUE: 0,
	}

	ct_cmyk := starting_color.To_cmyk()
	ct_hsv := starting_color.To_hsv()
	ct_hsl := starting_color.To_hsl()
	ct_lab := starting_color.To_cielab()
	ct_hex := starting_color.To_hex()

	res_container := color_container{
		rgb: &starting_color,
		cmyk: &ct_cmyk,
		hsv: &ct_hsv,
		hsl: &ct_hsl,
		lab: &ct_lab,
		hex: ct_hex,
	}

	return &res_container
}

func (c *color_container) update_values(modified_obj any) {
	switch v := modified_obj.(type) {
		case *colorspace.RGB_obj :
			obj := *v

			l_cmyk := obj.To_cmyk()
			l_hsv := obj.To_hsv()
			l_hsl := obj.To_hsl()
			l_lab := obj.To_cielab()

			c.cmyk = &l_cmyk
			c.hsv = &l_hsv
			c.hsl = &l_hsl
			c.lab = &l_lab
			c.hex = obj.To_hex()
		case *colorspace.CMYK_obj :
			obj := *v

			l_rgb := obj.To_rgb()
			l_hsv := obj.To_hsv()
			l_hsl := obj.To_hsl()
			l_lab := obj.To_cielab()

			c.rgb = &l_rgb
			c.hsv = &l_hsv
			c.hsl = &l_hsl
			c.lab = &l_lab
			c.hex = c.rgb.To_hex()
		case *colorspace.HSV_obj :
			obj := *v

			l_rgb := obj.To_rgb()
			l_cmyk := obj.To_cmyk()
			l_hsl := obj.To_hsl()
			l_lab := obj.To_cielab()
		
			c.rgb = &l_rgb
			c.cmyk = &l_cmyk
			c.hsl = &l_hsl
			c.lab = &l_lab
			c.hex = c.rgb.To_hex()
		case *colorspace.HSL_obj :
			obj := *v

			l_rgb := obj.To_rgb()
			l_cmyk := obj.To_cmyk()
			l_hsv := obj.To_hsv()
			l_lab := obj.To_cielab()

			c.rgb = &l_rgb
			c.cmyk = &l_cmyk
			c.hsv = &l_hsv
			c.lab = &l_lab
			c.hex = c.rgb.To_hex()
		case *colorspace.CIELAB_obj :
			obj := *v

			l_rgb := obj.To_rgb()
			l_cmyk := obj.To_cmyk()
			l_hsv := obj.To_hsv()
			l_hsl := obj.To_hsl()

			c.rgb = &l_rgb
			c.cmyk = &l_cmyk
			c.hsv = &l_hsv
			c.hsl = &l_hsl
			c.hex = c.rgb.To_hex()
		case string :
			l_rgb, _ := colorspace.Hex_to_rgb(v)
			l_cmyk := l_rgb.To_cmyk()
			l_hsv := l_rgb.To_hsv()
			l_hsl := l_rgb.To_hsl()
			l_lab := l_rgb.To_cielab()

			c.rgb = &l_rgb
			c.cmyk = &l_cmyk
			c.hsv = &l_hsv
			c.hsl = &l_hsl
			c.lab = &l_lab
		default :
			fmt.Println("Unknown datatype - update date", v)
	}
}

func bind_helper(ptr any) binding.Float {
	bind := binding.NewFloat()

	switch v := ptr.(type) {
		case *float64 :
			bind.Set(*v)
			bind.AddListener(binding.NewDataListener(func() {
				val, _ := bind.Get()
				*v = val
			}))
		case *uint8 :
			bind.Set(float64(*v))
			bind.AddListener(binding.NewDataListener(func() {
				val, _ := bind.Get()
				*v = uint8(val)
			}))
		default :
			fmt.Println("Unknown datatype - bind helper", ptr)
			return nil
	}

	return bind
}

func new_slider_field(kv any, mn, mx, step float64) *fyne.Container {
	var name string
	var value_ptr any

	switch v := kv.(type) {
		case map[string](*float64) :
			for i, val := range v {
				name = i
				value_ptr = val
			}
		case map[string](*uint8) :
			for i, val := range v {
				name = i
				value_ptr = val
			}
	}

	new_bind := bind_helper(value_ptr)

	text := widget.NewLabel(name)
	text.Resize(fyne.NewSize(15, 35))
	text.Move(fyne.NewPos(0,0))

	entry := widget.NewEntryWithData(binding.FloatToStringWithFormat(new_bind, "%.2f"))

	entry.MultiLine = false
	entry.Resize(fyne.NewSize(60,35))
	entry.Move(fyne.NewPos(330, 0))

	slider := widget.NewSliderWithData(mn, mx, new_bind)
	slider.Step = step

	slider.Resize(fyne.NewSize(315, 35))
	slider.Move(fyne.NewPos(20, 0))

	new_container := container.NewWithoutLayout(text, slider, entry)

	return new_container
}

func new_accordion(acc_name string, field_names []any, field_ranges [][]float64) *fyne.Container {
	sub_items := make([](*fyne.Container), len(field_names))

	for i, v := range field_names {
		sub_items[i] = new_slider_field(v, field_ranges[i][0], field_ranges[i][1], field_ranges[i][2])
	} 

	new_sub := make([]fyne.CanvasObject, len(sub_items))

	for i, v := range sub_items {
		new_sub[i] = v
	}

	sub_box := container.NewVBox(new_sub...)

	item := widget.NewAccordionItem(acc_name, sub_box)

	accordion := widget.NewAccordion(item)

	return container.NewVBox(accordion)
}

func new_hex_elem() *fyne.Container {
	hex_rect := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	hex_rect.SetMinSize(fyne.NewSize(120, 100))
	hex_rect.Resize(fyne.NewSize(120, 100))
	
	outline_rect := canvas.NewRectangle(color.RGBA{R:0, G: 0, B: 0, A: 255})
	outline_rect.SetMinSize(fyne.NewSize(131, 116))
	outline_rect.Resize(fyne.NewSize(131, 126))

	some_text := canvas.NewText("Lorem Ipsum...", color.RGBA{R: 0, G: 0, B: 0, A: 255})
	some_text.SetMinSize(fyne.NewSize(90, 20))
	some_text.Resize(fyne.NewSize(90, 20))

	hex_box := container.NewCenter(outline_rect, hex_rect, some_text)
	hex_box.Resize(fyne.NewSize(400, 116))

	entry := widget.NewEntry()
	entry.Resize(fyne.NewSize(80,35))
	entry.Move(fyne.NewPos(60, 21))

	entry_box := container.NewWithoutLayout(entry)

	check := widget.NewCheck("Bind text color to bg", func (checked bool) {
		fmt.Println("Some text")
	})

	check_box := container.NewCenter(check)

	final_box := container.NewGridWithColumns(2, hex_box, container.NewGridWithRows(2, entry_box, check_box))

	return final_box
}

func form_final_fields(hex_box *fyne.Container, cc *color_container) []fyne.CanvasObject {
	acc_names := []string{"CMYK", "HSV", "HSL", "CIE L*a*b* (CIELAB)"}

	rgb_fields := []any{map[string](*uint8){"R": &(cc.rgb.RED)}, map[string](*uint8){"G": &(cc.rgb.GREEN)}, map[string](*uint8){"B": &(cc.rgb.BLUE)}}

	rgb_field_limits := [][]float64{{0.0, 255.0, 1.0}, {0.0, 255.0, 1.0}, {0.0, 255.0, 1.0}}

	acc_field_names := [][]any{
		{map[string](*float64){"C": &(cc.cmyk.CYAN)}, map[string](*float64){"M": &(cc.cmyk.MAGENTA)}, map[string](*float64){"Y": &(cc.cmyk.YELLOW)}, map[string](*float64){"K": &(cc.cmyk.KEY)}},
		{map[string](*float64){"H": &(cc.hsv.HUE)}, map[string](*float64){"S": &(cc.hsv.SATURATION)}, map[string](*float64){"V": &(cc.hsv.VALUE)}},
		{map[string](*float64){"H": &(cc.hsl.HUE)}, map[string](*float64){"S": &(cc.hsl.SATURATION)}, map[string](*float64){"L": &(cc.hsl.LIGHTNESS)}},
		{map[string](*float64){"L": &(cc.lab.L)}, map[string](*float64){"a": &(cc.lab.A)}, map[string](*float64){"b": &(cc.lab.B)}},
	}

	acc_field_limits := [][][]float64{
		{{0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}},
		{{0.0, 360.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}},
		{{0.0, 360.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}},
		{{0.0, 100.0, 0.01}, {-150.0, 150.0, 0.01}, {-150.0, 150.0, 0.01}},
	}

	sub_acc_fields := make([]fyne.CanvasObject, (len(acc_names)*2) + 6)

	spacer_rect := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 0, A: 0})
	spacer_rect.SetMinSize(fyne.NewSize(400, 10))
	spacer_rect.Resize(fyne.NewSize(400, 10))

	sub_acc_fields[0] = spacer_rect
	sub_acc_fields[1] = hex_box
	sub_acc_fields[2] = spacer_rect
	sub_acc_fields[3] = widget.NewSeparator()

	new_acc := new_accordion("RGB (sRGB / Regular RGB)", rgb_fields, rgb_field_limits)

	sub_acc_fields[4] = new_acc
	sub_acc_fields[5] = widget.NewSeparator()

	for i, v := range acc_names {
		sub_acc_fields[(i*2) + 6] = new_accordion(v, acc_field_names[i], acc_field_limits[i])
		sub_acc_fields[(i*2) + 7] = widget.NewSeparator()
	}

	return sub_acc_fields
}

func main() {
	wrap_color := create_empty_container()

	wrap_color.rgb.RED += 10

	wrap_color.update_values(wrap_color.rgb)

	a := app.New()

	if img, err := fyne.LoadResourceFromPath("../assets/coolskull.png"); err == nil {
		a.SetIcon(img)
	}

	w := a.NewWindow("Color Picker")

	hex_elem := new_hex_elem()

	fields := form_final_fields(hex_elem, wrap_color)

	new_box := container.NewVBox(fields...)
	final_box := container.NewVScroll(new_box)

	bg := canvas.NewRectangle(color.RGBA{R: 35, G: 35, B: 35, A: 255})

	content := container.NewStack(bg, final_box)

	w.Resize(fyne.NewSize(400,400))
	w.SetFixedSize(true)
	w.SetContent(content)
	w.ShowAndRun()

}
