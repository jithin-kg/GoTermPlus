package terminal

import (
	"fyne.io/fyne/v2/widget"
)

type TerminalOutput struct {
	Output *widget.Entry
}

// func NewTerminalOuput() *TerminalOutput {
// 	// outputEntry := widget.NewMultiLineEntry()
// 	// outputEntry := NewCustomEntry()
// 	// outputEntry.MultiLine = true
// 	// outputEntry.Wrapping = fyne.TextWrapOff
// 	// outputEntry.Disable() //make it read only

// 	// styling
// 	return &TerminalOutput{
// 		Output: &outputEntry.Entry,
// 	}
// }

func (to *TerminalOutput) GetWidget() *widget.Entry {
	return to.Output
}

func (to *TerminalOutput) AppendText(text string) {
	to.Output.SetText(to.Output.Text + "\n" + text)
}
