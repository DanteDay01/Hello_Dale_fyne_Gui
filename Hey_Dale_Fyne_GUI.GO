package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hey Dale!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hello!", func() {
			hello.SetText("Good to see you :)")
		}),
	))

	w.ShowAndRun()
}
