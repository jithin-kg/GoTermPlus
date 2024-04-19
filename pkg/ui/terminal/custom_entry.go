package terminal

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomMultiLineEntry struct {
	widget.BaseWidget
	textEntry       *widget.Entry
	originalText    string
	scrollContainer *container.Scroll
}

func NewCustomMultiLineEntry() *CustomMultiLineEntry {
	e := &CustomMultiLineEntry{}
	e.ExtendBaseWidget(e)
	e.textEntry = widget.NewMultiLineEntry()
	e.textEntry.Wrapping = fyne.TextWrapBreak
	e.textEntry.TextStyle.Monospace = true
	e.scrollContainer = container.NewVScroll(e.textEntry)
	return e
}

func (e *CustomMultiLineEntry) SetText(text string) {
	e.originalText = text
	e.textEntry.SetText(text)
}

func (e *CustomMultiLineEntry) CreateRenderer() fyne.WidgetRenderer {
	return &customMultiLineEntryRenderer{entry: e, scrollContainer: e.scrollContainer}
}

type customMultiLineEntryRenderer struct {
	entry           *CustomMultiLineEntry
	scrollContainer *container.Scroll
}

func (r *customMultiLineEntryRenderer) MinSize() fyne.Size {
	return r.scrollContainer.MinSize()
}

func (r *customMultiLineEntryRenderer) Layout(size fyne.Size) {
	r.scrollContainer.Resize(size)
}

func (r *customMultiLineEntryRenderer) Refresh() {
	r.entry.textEntry.Refresh()
	r.scrollContainer.Refresh()
}

func (r *customMultiLineEntryRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.scrollContainer}
}

func (r *customMultiLineEntryRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *customMultiLineEntryRenderer) Destroy() {}

func (e *CustomMultiLineEntry) AppendText(text string) {
	e.originalText += text + "\n"
	e.textEntry.SetText(e.originalText)
	e.Refresh()
}

func (e *CustomMultiLineEntry) Refresh() {
	e.textEntry.Refresh()
	e.scrollContainer.Refresh()
	e.forceScrollToEnd()
}

func (e *CustomMultiLineEntry) forceScrollToEnd() {
	// Optional: Debugging print
	fmt.Println("Attempting to force scroll to end")
	e.scrollContainer.Refresh()

	time.AfterFunc(10*time.Millisecond, func() {
		scrollSize := e.scrollContainer.Content.Size().Height
		viewportSize := e.scrollContainer.Size().Height
		// if scrollSize > viewportSize {
		delta := scrollSize - viewportSize
		e.scrollContainer.Offset = fyne.NewPos(0, delta)
		e.scrollContainer.Refresh()
		e.scrollContainer.ScrollToBottom()
		fmt.Println("Scrolled to bottom: Delta", delta)
		// } else {
		// 	fmt.Println("No need to scroll: Content fits within viewport")
		// }
	})
}
