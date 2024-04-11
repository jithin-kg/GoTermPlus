package sshclient

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Client *ssh.Client
}

func NewSSHClient(client *ssh.Client) *SSHClient {
	return &SSHClient{Client: client}
}

func (c *SSHClient) ExecuteCommand(command string) (string, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()
	ouput, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("failed to execute command '%s' : %w", command, err)
	}
	return string(ouput), nil
}
func (c *SSHClient) ListDirectories(path string) ([]string, error) {
	output, err := c.ExecuteCommand(fmt.Sprintf("ls -l %s", path))
	if err != nil {
		return nil, err
	}
	lines := strings.Split(output, "\n")

	return lines, nil
}
