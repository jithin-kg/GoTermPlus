package custom

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// https://github.com/fyne-io/fyne/issues/3590
// ConsoleWriter encapsulates a VBox and a Scroll Container for console-like text output.
type ConsoleWriter struct {
	Console       *fyne.Container
	ConsoleScroll *container.Scroll
}

// NewConsoleWriter creates and returns a new ConsoleWriter component.
func NewConsoleWriter() *ConsoleWriter {
	vbox := container.NewVBox()
	scroll := container.NewVScroll(vbox)
	return &ConsoleWriter{
		Console:       vbox,
		ConsoleScroll: scroll,
	}
}

// Write adds text to the console and ensures it scrolls appropriately.
func (cw *ConsoleWriter) Write(text string) {
	entry := widget.NewLabel(text)
	// entry := &widget.Label{Text: text, Wrapping: fyne.TextWrapBreak, TextStyle:fyne.TextStyle{Monospace: true} }
	entry.Wrapping = fyne.TextWrapBreak // Ensure text wraps if too long
	entry.TextStyle = fyne.TextStyle{Monospace: true}

	cw.Console.Add(entry)
	cw.Console.Add(entry)
	// cw.Console.Add(&canvas.Text{
	// 	Text:      text,
	// 	Color:     color.White,
	// 	TextSize:  12,
	// 	TextStyle: fyne.TextStyle{Monospace: true},
	// })
	// This ensures that the UI updates occur after other operations have completed.
	time.AfterFunc(10*time.Millisecond, func() {
		// delta := (cw.Console.MinSize().Height - cw.ConsoleScroll.Size().Height) - cw.ConsoleScroll.Offset.Y
		// if delta < 50 {
		cw.ConsoleScroll.ScrollToBottom()
		// }
		cw.Console.Refresh()
	})
}

func (cw *ConsoleWriter) Scroll() {
	// cw.ConsoleScroll.ScrollToBottom()
}
