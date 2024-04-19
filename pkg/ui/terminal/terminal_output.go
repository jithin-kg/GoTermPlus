package terminal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TerminalOutput struct {
	Output *widget.Entry
}

func NewTerminalOuput() *TerminalOutput {
	outputEntry := widget.NewMultiLineEntry()
	outputEntry.MultiLine = true
	outputEntry.Wrapping = fyne.TextWrapOff
	outputEntry.Disable() //make it read only
	return &TerminalOutput{
		Output: outputEntry,
	}
}

func (to *TerminalOutput) GetWidget() *widget.Entry {
	return to.Output
}

func (to *TerminalOutput) AppendText(text string) {
	to.Output.SetText(to.Output.Text + "\n" + text)
}
