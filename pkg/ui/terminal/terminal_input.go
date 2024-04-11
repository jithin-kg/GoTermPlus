package terminal

import (
	"log"

	"github.com/jithin-kg/GoTermPlus/pkg/sshclient"
	custom "github.com/jithin-kg/GoTermPlus/pkg/ui/custom_widgets"
)

type TerminalInput struct {
	entry           *custom.MultiLineEntry
	client          *sshclient.SSHClient
	onCommandOutput func(string) //callback for handling command output
}

func NewTerminalInput(client *sshclient.SSHClient, onCommandOutput func(string)) *TerminalInput {
	ti := &TerminalInput{
		client:          client,
		entry:           custom.NewMultiLineEntry(),
		onCommandOutput: onCommandOutput,
	}
	ti.Setup()
	return ti
}
func (ti *TerminalInput) Setup() {
	ti.entry.SetPlaceHolder("Type command here...")
	ti.entry.OnSubmitted = func(content string) {
		if content != "" {
			go ti.executeCommand(content)
		}
	}
}

func (ti *TerminalInput) executeCommand(command string) {
	output, err := ti.client.ExecuteCommand(command)
	if err != nil {
		log.Printf("Error executing command '%s': %v", command, err)
		output = "Error: " + err.Error()
	}
	if ti.onCommandOutput != nil {
		ti.onCommandOutput(command + "\n" + output)
	}
	ti.entry.AddHistory(command)
	// terminalOutput.SetText(terminalInput.Text + command + "\n" + output + "\n")
	ti.entry.SetText("") // clear input field

}

// returns the entry widget for use in UI layouts
func (ti *TerminalInput) GetEtry() *custom.MultiLineEntry {
	return ti.entry
}
