package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Diwice/color-picker/pkg/colorspace"
	)

func main() {
	my_color := colorspace.RGB_obj{
		RED: 255,
		GREEN: 0,
		BLUE: 0,
	}

	my_app := app.New()
	my_window := my_app.NewWindow("Hello")
	my_window.SetContent(widget.NewLabel(fmt.Sprintf("The color you've made : R-%d/G-%d/B-%d", my_color.RED, my_color.GREEN, my_color.BLUE))

	my_window.Show()
	my_app.Run()
}
