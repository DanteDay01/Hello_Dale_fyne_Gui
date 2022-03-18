package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Print("Hey Dale Fyne GUI")
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hey Dale!")
	containerTimes := container.NewVBox() // One item sized with scroll bar
	//containerTimes := container.NewPadded()	// Multiple items only on-top of each other.
	//containerTimes := container.NewMax()	// One item only
	container1 := container.NewVBox(
		hello,
		widget.NewButton("Hello!", func() {
			// Format dates https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html
			// Date here, so I see a different result. Or perhaps stream The messages one after the other... Can I add
			// another text box above this button?
			now := time.Now().Format("2006.01.02 15:04:05")
			containerTimes.Add(widget.NewLabel(now))
			//hello.SetText("Good to see you " + now)
			//hello.SetText("Good to see you " + now.String())
			//containerTimes.Refresh()	// No obvious effect. Can I make the Refre() happen when container1 is resized?
		}),
		// TODO, can I set the NewVScroll to grow to grow with the container it's inside?
		// 		https://developer.fyne.io/faq/layout
		//		https://developer.fyne.io/extend/custom-layout
		//		https://stackoverflow.com/questions/69685039/scrolling-log-text-in-a-fyne-textgrid-automatically
		//		https://github.com/fyne-io/fyne/issues/259
		//		https://stackoverflow.com/questions/68005561/how-to-resize-an-entry-box-in-fyne
		//		https://developer.fyne.io/binding/list.html

		//widget.NewLabel("VScroll:"),
		container.NewVScroll(containerTimes),
		//widget.NewLabel("VBox:"),
		//containerTimes,
		//container.NewVScroll(containerTimes),
	)
	w.SetContent(container1)
	/*
		w.SetContent(container.NewVBox(
			hello,
			widget.NewButton("Hello!", func() {
				// Add something here (perhaps the date and time) so I see a different result. Or perhaps stream
				//		The messages one after the other... Can I add another text box above this button?
				hello.SetText("Good to see you :)")
			}),
		))
	*/

	w.ShowAndRun()
}
