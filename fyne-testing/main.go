package main

import (
	"fmt"
	"math"
	"strconv"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
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

func (c *color_container) up_rgb(obj colorspace.RGB_obj) {
	l_cmyk := obj.To_cmyk()
	l_hsv := obj.To_hsv()
	l_hsl := obj.To_hsl()
	l_lab := obj.To_cielab()

	c.cmyk = &l_cmyk
	c.hsv = &l_hsv
	c.hsl = &l_hsl
	c.lab = &l_lab
	c.hex = obj.To_hex()
}

func (c *color_container) up_cmyk(obj colorspace.CMYK_obj) {
	l_rgb := obj.To_rgb()
	l_hsv := obj.To_hsv()
	l_hsl := obj.To_hsl()
	l_lab := obj.To_cielab()

	c.rgb = &l_rgb
	c.hsv = &l_hsv
	c.hsl = &l_hsl
	c.lab = &l_lab
	c.hex = c.rgb.To_hex()
}

func (c *color_container) up_hsv(obj colorspace.HSV_obj) {
	l_rgb := obj.To_rgb()
	l_cmyk := obj.To_cmyk()
	l_hsl := obj.To_hsl()
	l_lab := obj.To_cielab()

	c.rgb = &l_rgb
	c.cmyk = &l_cmyk
	c.hsl = &l_hsl
	c.lab = &l_lab
	c.hex = c.rgb.To_hex()
}

func (c *color_container) up_hsl(obj colorspace.HSL_obj) {
	l_rgb := obj.To_rgb()
	l_cmyk := obj.To_cmyk()
	l_hsv := obj.To_hsv()
	l_lab := obj.To_cielab()

	c.rgb = &l_rgb
	c.cmyk = &l_cmyk
	c.hsv = &l_hsv
	c.lab = &l_lab
	c.hex = c.rgb.To_hex()
}

func (c *color_container) up_lab(obj colorspace.CIELAB_obj) {
	l_rgb := obj.To_rgb()
	l_cmyk := obj.To_cmyk()
	l_hsv := obj.To_hsv()
	l_hsl := obj.To_hsl()

	c.rgb = &l_rgb
	c.cmyk = &l_cmyk
	c.hsv = &l_hsv
	c.hsl = &l_hsl
	c.hex = c.rgb.To_hex()
}

func (c *color_container) up_hex(obj string) {
	l_rgb, _ := colorspace.Hex_to_rgb(obj)
	l_cmyk := l_rgb.To_cmyk()
	l_hsv := l_rgb.To_hsv()
	l_hsl := l_rgb.To_hsl()
	l_lab := l_rgb.To_cielab()

	c.rgb = &l_rgb
	c.cmyk = &l_cmyk
	c.hsv = &l_hsv
	c.hsl = &l_hsl
	c.lab = &l_lab
}

func (c *color_container) update_values(modified_obj any) {
	switch modified_obj.(type) {
		case *colorspace.RGB_obj :
			ptr, _ := modified_obj.(*colorspace.RGB_obj)
			c.up_rgb(*ptr)
		case *colorspace.CMYK_obj :
			ptr, _ := modified_obj.(*colorspace.CMYK_obj)
			c.up_cmyk(*ptr)
		case *colorspace.HSV_obj :
			ptr, _ := modified_obj.(*colorspace.HSV_obj)
			c.up_hsv(*ptr)
		case *colorspace.HSL_obj :
			ptr, _ := modified_obj.(*colorspace.HSL_obj)
			c.up_hsl(*ptr)
		case *colorspace.CIELAB_obj :
			ptr, _ := modified_obj.(*colorspace.CIELAB_obj)
			c.up_lab(*ptr)
		case string :
			c.up_hex(modified_obj.(string))
		default :
			fmt.Println("Unknown datatype")
	}
}

type custom_rect struct {
	*canvas.Rectangle
	width, height float32
}

func (o *custom_rect) MinSize() fyne.Size {
	return fyne.NewSize(o.width, o.height)
}

func (o *custom_rect) Resize(size fyne.Size) {
	o.Rectangle.Resize(size)
}

func (o *custom_rect) Move(pos fyne.Position) {
	o.Rectangle.Move(pos)
}

func (o *custom_rect) Show() {
	o.Rectangle.Show()
}

func (o *custom_rect) Hide() {
	o.Rectangle.Hide()
}

func (o *custom_rect) Refresh() {
	o.Rectangle.Refresh()
}

func new_hex_rect(w, h float32) *custom_rect {
	new_rect := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	new_rect.Resize(fyne.NewSize(w, h))

	res_rect := &custom_rect{
		Rectangle: new_rect,
		width: w,
		height: h,
	}

	return res_rect
}

func new_slider_field(name string, mn, mx, step float64) *fyne.Container {
	text := widget.NewLabel(name)
	text.Resize(fyne.NewSize(15, 35))
	text.Move(fyne.NewPos(0,0))

	entry := widget.NewEntry()

	if math.Mod(step, 1.0) != 0.0 {
		entry.SetText("0.00")
	} else {
		entry.SetText("0")
	}

	entry.MultiLine = false
	
	entry.Resize(fyne.NewSize(60,35))
	entry.Move(fyne.NewPos(330, 0))

	slider := widget.NewSlider(mn, mx)
	slider.Step = step

	slider.Resize(fyne.NewSize(315, 35))
	slider.Move(fyne.NewPos(20, 0))

	entry.OnSubmitted = func(text string) {
		if v, err := strconv.ParseFloat(text, 64); err == nil {
			if v >= slider.Min && v <= slider.Max {
				slider.SetValue(v)
			} else if v < slider.Min {
				slider.SetValue(slider.Min)
			} else {
				slider.SetValue(slider.Max)
			}
		}
	}

	slider.OnChanged = func(val float64) {
		if math.Mod(val, 1.0) != 0.0 {
			entry.SetText(fmt.Sprintf("%.2f", val))
		} else {
			entry.SetText(fmt.Sprintf("%d", int(val)))
		}
	}

	new_container := container.NewWithoutLayout(text, slider, entry)

	return new_container
}

func new_accordion(acc_name string, field_names []string, field_ranges [][]float64) *fyne.Container {
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

func main() {
	wrap_color := create_empty_container()

	wrap_color.rgb.RED += 10

	a := app.New()

	if img, err := fyne.LoadResourceFromPath("../assets/coolskull.png"); err == nil {
		a.SetIcon(img)
	}

	w := a.NewWindow("Color Picker")

	hex_rect := new_hex_rect(100.0, 50.0)
	hex_box := container.NewMax(hex_rect)
	hex_box.Resize(fyne.NewSize(hex_rect.width, hex_rect.height))

	acc_names := []string{"RGB (sRGB / Regular RGB)", "CMYK", "HSV", "HSL", "CIE L*a*b* (CIELAB)"}

	acc_field_names := [][]string{
		{"R", "G", "B"},
		{"C", "M", "Y", "K"},
		{"H", "S", "V"},
		{"H", "S", "L"},
		{"L", "a", "b"},
	}

	acc_field_limits := [][][]float64{
		{{0.0, 255.0, 1.0}, {0.0, 255.0, 1.0}, {0.0, 255.0, 1.0}},
		{{0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}},
		{{0.0, 360.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}},
		{{0.0, 360.0, 0.01}, {0.0, 100.0, 0.01}, {0.0, 100.0, 0.01}},
		{{0.0, 100.0, 0.01}, {-150.0, 150.0, 0.01}, {-150.0, 150.0, 0.01}},
	}

	acc_fields := make([](*fyne.Container), len(acc_names))
	sub_acc_fields := make([]fyne.CanvasObject, len(acc_fields)+1)

	sub_acc_fields[0] = hex_box

	for i, v := range acc_names {
		acc_fields[i] = new_accordion(v, acc_field_names[i], acc_field_limits[i])
		sub_acc_fields[i+1] = acc_fields[i]
	}

	//final_box := container.NewVScroll(container.NewVBox(sub_acc_fields...))
	final_box := container.NewVScroll(container.NewVBox(sub_acc_fields...))
	bg := canvas.NewRectangle(color.RGBA{R: 35, G: 35, B: 35, A: 255})

	content := container.NewMax(bg, final_box)

	w.Resize(fyne.NewSize(400,400))
	w.SetFixedSize(true)
	w.SetContent(content)
	w.ShowAndRun()
}
