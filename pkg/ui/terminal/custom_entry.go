package terminal

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// CustomMultiLineEntry struct holds the text content and the scroll container.
type CustomMultiLineEntry struct {
	widget.BaseWidget
	textLabel       *widget.Label
	scrollContainer *container.Scroll
}

// NewCustomMultiLineEntry initializes the CustomMultiLineEntry with default settings.
func NewCustomMultiLineEntry() *CustomMultiLineEntry {
	e := &CustomMultiLineEntry{}
	e.ExtendBaseWidget(e)
	e.textLabel = widget.NewLabel("")
	e.textLabel.Wrapping = fyne.TextWrapBreak // Enable word wrapping
	e.textLabel.TextStyle.Monospace = true    // Optional: Set monospace for terminal style
	e.scrollContainer = container.NewVScroll(container.NewWithoutLayout(e.textLabel))
	return e
}

// CreateRenderer creates a renderer for CustomMultiLineEntry.
func (e *CustomMultiLineEntry) CreateRenderer() fyne.WidgetRenderer {
	return &customMultiLineEntryRenderer{entry: e, scrollContainer: e.scrollContainer}
}

// Custom renderer for CustomMultiLineEntry.
type customMultiLineEntryRenderer struct {
	entry *CustomMultiLineEntry
	// textContent     *canvas.Text
	scrollContainer *container.Scroll
}

// MinSize calculates the minimum size of the widget.
func (r *customMultiLineEntryRenderer) MinSize() fyne.Size {
	return r.scrollContainer.MinSize()
}

// Layout resizes and positions the widget's elements.
func (r *customMultiLineEntryRenderer) Layout(size fyne.Size) {
	r.scrollContainer.Resize(size)
}

// Refresh updates the widget when the data changes.
func (r *customMultiLineEntryRenderer) Refresh() {
	// r.textContent.Refresh()
	r.entry.textLabel.Refresh()
	r.scrollContainer.Refresh()
}

// Objects returns all objects in the renderer.
func (r *customMultiLineEntryRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.scrollContainer}
}

// BackgroundColor returns the background color of the widget.
func (r *customMultiLineEntryRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

// Destroy cleans up any resources.
func (r *customMultiLineEntryRenderer) Destroy() {}

// AppendText adds new text to the existing content.
func (e *CustomMultiLineEntry) AppendText(text string) {
	// e.textContent.Text += text + "\n" // Append new text as a new line.
	e.textLabel.SetText(e.textLabel.Text + text + "\n") // Append new text as a new line.
	e.Refresh()
}

// Refresh updates the text and scrolls to the bottom.
func (e *CustomMultiLineEntry) Refresh() {
	e.BaseWidget.Refresh()
	e.scrollContainer.ScrollToBottom() // Auto-scroll to bottom.
}
