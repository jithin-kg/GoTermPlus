package terminal

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jithin-kg/GoTermPlus/pkg/sshclient"
)

func NewTerminalScreen(window fyne.Window, client *sshclient.SSHClient) fyne.CanvasObject {

	data := binding.BindStringList(&[]string{})
	// Create the file browser pane using NewListWithData and the data binding
	fileBrowserScroll := createFileBrowser(data)
	terminalScroll := createTerminalPane(client)
	topPanel := createTopPanel()

	// assemble screen using horizontal split container
	hSplit := container.NewHSplit(fileBrowserScroll, terminalScroll)
	hSplit.Offset = 0.25 // initial split ratio
	content := container.New(NewLastItemFullheightVBoxLayout(), topPanel, hSplit)

	// perform directory listing in a seperate goroutin
	populateFileBrowserAsync(data, client)
	return content
}

func populateFileBrowserAsync(data binding.StringList, client *sshclient.SSHClient) {
	go func() {
		listItems, err := client.ListDirectories("/")
		if err != nil {
			log.Printf("Failed to list directories %v", err)
		}
		data.Set(listItems)

	}()
}

func createFileBrowser(data binding.StringList) *container.Scroll {
	fileBrowser := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			// this function defines how each item in list looks like
			// Return the template for list items
			return widget.NewLabel("")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			// for each item in the list this function get called providin the item's data (i)
			// Bind the list item to the data item
			str, _ := i.(binding.String).Get() // Extract the string from the binding
			o.(*widget.Label).SetText(str)     // Set the label text
		},
	)
	return container.NewScroll(fileBrowser)
}

func createTopPanel() *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Options:"),
		widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {}),
		// Add more toolbar actions as needed
		layout.NewSpacer(), // Pushes everything to the left
	)
}
func createTerminalPane(client *sshclient.SSHClient) *fyne.Container {
	// output area for command results
	terminalOutput := widget.NewMultiLineEntry()
	terminalOutput.MultiLine = true
	terminalOutput.Wrapping = fyne.TextWrapOff
	terminalOutput.Disable() //make it read only
	// terminalInput is for typing commands
	terminalInput := NewTerminalInput(client, func(output string) {
		terminalOutput.SetText(terminalOutput.Text + "\n" + output)
	})
	terminalLayout := container.NewBorder(nil, terminalInput.GetEtry(), nil, nil, terminalOutput)

	return container.NewPadded(terminalLayout)
}
