package terminal

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/ssh"
)

func NewTerminalScreen(window fyne.Window, client *ssh.Client) fyne.CanvasObject {
	// list directories and files
	listItems, err := listDirectories(client, "/")
	if err != nil {
		log.Printf("Failed to list directories %v", err)
	}

	// file browser pane
	fileBrowser := setupFileBrowser(listItems)
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
	fmt.Println(topPanel)
	// assemble screen using horizontal split container
	hSplit := container.NewHSplit(fileBrowserScroll, terminalScroll)
	hSplit.Offset = 0.25 // initial split ratio
	// hSplit.StretchRight
	// return container.New(layout.NewGridLayout(2), fileBrowserScroll, terminalScroll)
	return container.NewBorder(topPanel, nil, nil, hSplit)
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
func listDirectories(client *ssh.Client, path string) ([]string, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()
	// command to list directories and files
	cmd := fmt.Sprintf("ls -l %s", path)
	ouput, err := session.CombinedOutput(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command %s: %w", cmd, err)
	}
	lines := strings.Split(string(ouput), "\n")
	return lines, nil
}
