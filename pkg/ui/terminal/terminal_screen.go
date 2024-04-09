package terminal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewTerminalScreen(window fyne.Window) fyne.CanvasObject {
	// file browser pane
	fileBrowser := widget.NewList(func() int {
		// returns the number of items in the list
		return 0
	},
		func() fyne.CanvasObject {
			// return the list item UI
			return nil
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			// update the list item UI
		},
	)
	// create terminal pane
	terminal := widget.NewMultiLineEntry()
	terminal.SetPlaceHolder("Terminal will be here")

	// top panel with options and settings
	topPanel := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			// implement copy functionality
		}),
		// additional toolbar items
	)

	// assemble screen using horizontal split container
	hSplit := container.NewHSplit(fileBrowser, terminal)
	hSplit.Offset = 0.3 // initial split ratio
	return container.NewBorder(topPanel, nil, nil, hSplit)
}
