package terminal

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// CustomMultiLineEntry struct holds the text content and the scroll container.
type CustomMultiLineEntry struct {
	widget.BaseWidget
	textEntry       *widget.Entry
	originalText    string // Store the original text
	scrollContainer *container.Scroll
}

// NewCustomMultiLineEntry initializes the CustomMultiLineEntry with default settings.
func NewCustomMultiLineEntry() *CustomMultiLineEntry {
	e := &CustomMultiLineEntry{}
	e.ExtendBaseWidget(e)
	e.textEntry = widget.NewMultiLineEntry()
	e.textEntry.OnChanged = func(s string) {
		fmt.Print("on -changed")
		e.textEntry.SetText(e.originalText)
	}
	e.textEntry.Wrapping = fyne.TextWrapBreak
	e.textEntry.TextStyle.Monospace = true
	e.scrollContainer = container.NewVScroll(e.textEntry)
	return e
}

// SetText sets the text and stores it as the original text
func (e *CustomMultiLineEntry) SetText(text string) {
	e.originalText = text // Update original text
	e.textEntry.SetText(text)
}

// CreateRenderer creates a renderer for CustomMultiLineEntry.
func (e *CustomMultiLineEntry) CreateRenderer() fyne.WidgetRenderer {
	return &customMultiLineEntryRenderer{entry: e, scrollContainer: e.scrollContainer}
}

// Custom renderer for CustomMultiLineEntry.
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

// AppendText adds new text to the existing content.
func (e *CustomMultiLineEntry) AppendText(text string) {
	e.originalText += text + "\n" // Append new text as a new line
	e.textEntry.SetText(e.originalText)
	e.Refresh()
}

// Refresh updates the text and scrolls to the bottom.
func (e *CustomMultiLineEntry) Refresh() {
	e.BaseWidget.Refresh()
	e.scrollContainer.ScrollToBottom() // Auto-scroll to bottom.
}

// OnChanged is an overridden method to intercept changes and prevent them
func (e *CustomMultiLineEntry) OnChanged() {

	// Reset to original text to ignore any user changes
	e.textEntry.SetText(e.originalText)
}
