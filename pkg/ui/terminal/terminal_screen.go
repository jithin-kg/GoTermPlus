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

	// file browser pane
	// fileBrowser := setupFileBrowser(nil)
	data := binding.BindStringList(&[]string{})

	// Create the file browser pane using NewListWithData and the data binding
	fileBrowser := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			// Return the template for list items
			return widget.NewLabel("")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			// Bind the list item to the data item
			str, _ := i.(binding.String).Get() // Extract the string from the binding
			o.(*widget.Label).SetText(str)     // Set the label text
		},
	)
	fileBrowserScroll := container.NewScroll(fileBrowser)

	container.NewScroll(fileBrowser)
	// create terminal pane
	terminal := setupTerminal()
	terminalScroll := container.NewVScroll(terminal)
	terminalScroll.SetMinSize(fyne.NewSize(400, 300)) // Set a minimum size for the terminal pane

	// top panel with options and settings
	topPanel := container.NewHBox(
		widget.NewLabel("Options:"),
		widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {}),
		// Add more toolbar actions as needed
		layout.NewSpacer(), // Pushes everything to the left
	)
	// assemble screen using horizontal split container
	hSplit := container.NewHSplit(fileBrowserScroll, terminalScroll)
	hSplit.Offset = 0.25 // initial split ratio
	// return hSplit
	content := container.New(NewLastItemFullheightVBoxLayout(), topPanel, hSplit)

	// list directories and files
	// perform directory listing in a seperate goroutin
	go func() {
		listItems, err := client.ListDirectories("/")
		if err != nil {
			log.Printf("Failed to list directories %v", err)
		}
		data.Set(listItems)

	}()
	return content
}
func setupTerminal() *widget.Entry {
	// Create terminal pane with multiline entry
	terminal := widget.NewMultiLineEntry()
	terminal.MultiLine = true // Enable multiline
	terminal.SetPlaceHolder("Terminal will be here")

	return terminal
}
func setupFileBrowser(listItems []string) *widget.List {
	fileItem := widget.NewLabel("")
	fileItem.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)

	fb := widget.NewList(func() int {
		// returns the number of items in the list
		return len(listItems)
	},
		func() fyne.CanvasObject {
			// return the list item UI
			return container.NewHBox(fileItem, layout.NewSpacer())
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			// update the list item UI
			label := obj.(*fyne.Container).Objects[0].(*widget.Label)
			label.SetText(listItems[id]) //set the lable text to directory/ file name
		},
	)

	return fb
}
