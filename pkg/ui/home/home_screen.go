package home

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/ssh"
)

func NewHomeScreen(window fyne.Window) fyne.CanvasObject {
	// form items
	username := widget.NewEntry()
	password := widget.NewPasswordEntry()
	hostEntry := widget.NewEntry()

	sshKey := widget.NewButton("Select SSH key", func() {
		// key selection logic
	})
	connectBtn := widget.NewButton("Connect", func() {
		// Implement connection logic
		config := ssh.ClientConfig{
			User: username.Text,
			Auth: []ssh.AuthMethod{
				ssh.Password(password.Text),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), //NOTE: not recommented for production, just for testing
			Timeout:         5 * time.Second,
		}
		client, err := ssh.Dial("tcp", hostEntry.Text, &config)
		if err != nil {
			log.Println("Failed to dial: ", err)
			// window.Canvas().Refresh(formContainer) // update the ui to show any error message
			return
		}
		session, err := client.NewSession()
		if err != nil {
			log.Println("Failed to create session:", err)
			client.Close()
			// window.Canvas().Refresh(formContainer)
			return
		}
		defer session.Close()
		log.Println("SSH connection established")

	})

	//Create the form
	form := widget.NewForm(
		widget.NewFormItem("Host:port", hostEntry),
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
	)
	form.OnSubmit = func() {
		// Implement the form submission logic here
	}
	formContainer := container.NewVBox(form, sshKey, connectBtn)
	// return container with form and submit button
	return formContainer
}

func onConnectClick() {

}
