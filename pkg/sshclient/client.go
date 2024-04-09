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
func (c *SSHClient) ListDirectories(path string) ([]string, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()
	// command to list directories and files
	cmd := "ls -l ~"
	// cmd := "ls /home/go-code/sample"
	ouput, err := session.CombinedOutput(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command %s: %w", cmd, err)
	}

	lines := strings.Split(string(ouput), "\n")
	fmt.Println("ouput", ouput)
	fmt.Println("lines", lines)
	return lines, nil
}
