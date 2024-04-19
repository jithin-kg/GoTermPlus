package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Console = container.NewVBox()
var ConsoleScrollBar = container.NewVScroll(Console)

func ConsoleWriter(text string) {
	Console.Add(&canvas.Text{
		Text:      text,
		Color:     color.White,
		TextSize:  12,
		TextStyle: fyne.TextStyle{Monospace: true},
	})
	if len(Console.Objects) > 100 {
		Console.Remove(Console.Objects[0])
	}

	delta := (Console.MinSize().Height - ConsoleScrollBar.Size().Height) - ConsoleScrollBar.Offset.Y
	if delta < 50 {
		ConsoleScrollBar.ScrollToBottom()
	}
	Console.Refresh()
}

func main2() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Scroll to Bottom Example")
	myWindow.Resize(fyne.NewSize(400, 300))

	// Populate the console with initial text to ensure it needs scrolling
	for i := 0; i < 100; i++ {
		ConsoleWriter(fmt.Sprintf("This is line number %d", i))
	}

	scrollButton := widget.NewButton("Scroll to Bottom", func() {
		ConsoleScrollBar.ScrollToBottom()
	})
	content := container.NewBorder(nil, scrollButton, nil, nil, ConsoleScrollBar)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
