package terminal

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type scrollingEntry struct {
	widget.Entry
	scroller *container.Scroll
}

func (e *scrollingEntry) AppendText(text string) {
	// e.Entry.Text += text + "\n" // Append new text as a new line
	e.Entry.SetText(e.Entry.Text + text + "\n")
	e.Refresh()
	time.AfterFunc(10*time.Millisecond, func() {
		e.scroller.ScrollToBottom()
	})
}
func newScrollingEntry(scroller *container.Scroll) *scrollingEntry {
	entry := &scrollingEntry{}
	entry.ExtendBaseWidget(entry)
	entry.scroller = scroller
	return entry
}

func (e *scrollingEntry) TypedRune(r rune) {
	e.Entry.TypedRune(r)
	e.scheduleScroll()
}

func (e *scrollingEntry) TypedKey(key *fyne.KeyEvent) {
	e.Entry.TypedKey(key)
	e.scheduleScroll()
}

func (e *scrollingEntry) scheduleScroll() {
	time.AfterFunc(10*time.Millisecond, func() {
		e.scroller.ScrollToBottom()
	})
}
