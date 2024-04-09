package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/jithin-kg/GoTermPlus/pkg/ui/home"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GoTermPlus")

	// set initial content as home screen
	myWindow.SetContent(home.NewHomeScreen(myWindow))

	// set min width and height for window
	myWindow.Resize(fyne.NewSize(800, 600))
	//later switch to terminal screen like this:
	// myWindow.SetContent(terminal.NewTerminalScreen(myWindow))
	myWindow.ShowAndRun()
}
