package home

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jithin-kg/GoTermPlus/pkg/ui/terminal"
	"golang.org/x/crypto/ssh"
)

func NewHomeScreen(window fyne.Window) fyne.CanvasObject {
	// form items
	username := widget.NewEntry()
	username.SetText("root")
	password := widget.NewPasswordEntry()
	password.SetText("root")
	hostEntry := widget.NewEntry()
	hostEntry.SetText("localhost:2222")
	statusLabel := widget.NewLabel("")
	statusLabel.Wrapping = fyne.TextWrapWord

	sshKey := widget.NewButton("Select SSH key", func() {
		// key selection logic
	})

	connectBtn := widget.NewButton("Connect", func() {
		// connection logic
		config := CreateSShClientConfig(username.Text, password.Text)
		details := &SSHConnectionDetails{
			Config:  config,
			Address: hostEntry.Text,
			Status:  statusLabel,
		}
		details.OnSuccess = func(client *ssh.Client) {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Connection Established",
				Content: "Ssh connection established sucessfully.",
			})
			window.SetContent(terminal.NewTerminalScreen(window, client))
		}
		go connectToSSH(details)

	})

	//Create the form
	form := widget.NewForm(
		widget.NewFormItem("Host:port", hostEntry),
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Password", password),
	)
	form.OnSubmit = connectBtn.OnTapped //bind the form submission to the button's tap handler
	formContainer := container.NewVBox(form, sshKey, connectBtn, statusLabel)
	// return container with form and submit button
	return formContainer
}

func CreateSShClientConfig(username, password string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //NOTE: not recommented for production, just for testing
		Timeout:         5 * time.Second,
	}
}

type SSHConnectionDetails struct {
	Config    *ssh.ClientConfig
	Address   string
	Status    *widget.Label
	Client    *ssh.Client
	OnSuccess func(client *ssh.Client) // Callback function
}

func connectToSSH(details *SSHConnectionDetails) {
	client, err := ssh.Dial("tcp", details.Address, details.Config)
	if err != nil {
		log.Println("Failed to dial: ", err)
		details.Status.SetText(fmt.Sprintf("Failed to dial: %s", err))
		return
	}
	session, err := client.NewSession()
	if err != nil {
		log.Println("Failed to create session:", err)
		client.Close()
		details.Status.SetText(fmt.Sprintf("Failed to create session: %s", err))
		return
	}
	defer session.Close()
	details.Client = client
	if details.OnSuccess != nil {
		details.OnSuccess(client)
	}
	details.Status.SetText("SSH connection established")

	log.Println("SSH connection established")
}
