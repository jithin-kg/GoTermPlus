package custom

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// TerminalEntry extends widget.Entry to support command history

type MultiLineEntry struct {
	widget.Entry
	history       []string     // command history
	historyIndex  int          // current position of command history
	onHistoryFunc func(string) // function to call when a history item is selected

}

func NewMultiLineEntry() *MultiLineEntry {
	entry := &MultiLineEntry{}
	entry.MultiLine = true
	entry.ExtendBaseWidget(entry)
	return entry
}

// OnKeyTyped handles key events to navigate command history
func (t *MultiLineEntry) TypedKey(key *fyne.KeyEvent) {

	switch key.Name {
	case fyne.KeyUp:
		log.Println("TypedKey fyne.KeyUp", fyne.KeyUp)
		// enavigate up in the command history

		if t.historyIndex > 0 {
			t.historyIndex--
			t.SetText(t.history[t.historyIndex])
			if t.onHistoryFunc != nil {
				t.onHistoryFunc(t.history[t.historyIndex])
			}

		}
	case fyne.KeyDown:
		// navigat down in the command history

		if t.historyIndex < len(t.history)-1 {
			t.historyIndex++
			t.SetText(t.history[t.historyIndex])
			if t.onHistoryFunc != nil {
				t.onHistoryFunc(t.history[t.historyIndex])
			}

		} else if t.historyIndex == len(t.history)-1 {
			// clear the entry if we are at the end of history
			t.historyIndex = len(t.history)
			t.SetText("")
		}
	case fyne.KeyReturn:
		if t.Text != "" {
			if t.OnSubmitted != nil {
				t.OnSubmitted(t.Text)
			}
		}
	default:
		// call the parent handler for other keys
		t.Entry.TypedKey(key)
	}

}

func (t *MultiLineEntry) AddHistory(command string) {
	t.history = append(t.history, command)
	t.historyIndex = len(t.history)
}

func (t *MultiLineEntry) SetOnHistoryFunc(f func(string)) {
	t.onHistoryFunc = f
}
