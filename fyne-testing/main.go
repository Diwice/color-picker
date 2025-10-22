package main

import (
	"fmt"
	"strconv"
	//"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	//"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func new_slider_field(mn, mx float64) *fyne.Container {
	entry := widget.NewEntry()
	entry.SetText("0.00")

	slider := widget.NewSlider(mn, mx)
	slider.Step = 0.01

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
		entry.SetText(fmt.Sprintf("%.2f", val))
	}

	new_container := container.NewHBox(slider, layout.NewSpacer(), entry)

	return new_container
}

func main() {
	a := app.New()
	w := a.NewWindow("Testing")

	item_one := widget.NewAccordionItem("Item 1", widget.NewLabel("This sucks"))
	
	/*fixed_w := fyne.NewSize(60, entry.MinSize().Height)
	entry_box := container.New(layout.NewHBoxLayout(), entry)
	entry_box.Resize(fixed_w)
	entry_box = container.NewMax(entry_box) */

	sub_items := new_slider_field(0.00, 100.00)

	item_two := widget.NewAccordionItem("Item 2", sub_items)

	accordion := widget.NewAccordion(item_one, item_two)

	w.Resize(fyne.NewSize(400,150))
	w.SetContent(accordion)
	w.ShowAndRun()
}
