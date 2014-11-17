package utils

import (
	"code.google.com/p/go.crypto/ssh"
)

type SshClient interface {
	Close() error
	ExecCommand(command string) (string, error)
}

type softlayerSshClient struct {
	client *ssh.Client
}

func GetSshClient(username string, password string, ip string) (*softlayerSshClient, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	client, err := ssh.Dial("tcp", ip+":22", config)

	return &softlayerSshClient{client: client}, err
}

func (sshClient *softlayerSshClient) Close() error {
	return sshClient.client.Close()
}

func (sshClient *softlayerSshClient) ExecCommand(command string) (string, error) {
	session, err := sshClient.client.NewSession()
	if err != nil {
	}
	defer session.Close()

	output, err := session.Output(command)

	return string(output), err
}
