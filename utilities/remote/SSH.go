package remote

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Connection struct {
	Client *ssh.Client
	Mode   *ssh.TerminalModes
}

func makeConnection(conn *ssh.Client) Connection {
	return Connection{
		Client: conn,
		Mode:   &ssh.TerminalModes{ssh.ECHO: 0},
	}
}

func SSH(function func(Connection) error) error {
	config := &ssh.ClientConfig{
		User: parameters.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(parameters.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connection, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", parameters.Host, parameters.Port), config)
	if err != nil {
		return err
	}
	defer connection.Close()

	conn := makeConnection(connection)

	return function(conn)
}

func (c *Connection) getSession() (*ssh.Session, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		return session, err
	}

	err = session.RequestPty("xterm", 50, 80, *c.Mode)
	if err != nil {
		return session, err
	}
	return session, err
}

func (c *Connection) Shell(command string) (string, string, error) {
	session, err := c.getSession()
	if err != nil {
		return "", "", err
	}
	defer session.Close()

	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(command)

	return stdout.String(), stderr.String(), err
}

func (c *Connection) ShellOutput(command string) error {
	session, err := c.getSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if parameters.Output == "" {
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr

		err = session.Run(command)
	} else {
		var (
			out    bytes.Buffer
			report bytes.Buffer
		)

		writer := io.MultiWriter(os.Stdout, &out)

		session.Stdout = writer
		session.Stderr = os.Stderr

		err = session.Run(command)

		if strings.Contains(out.String(), "====== Report ======") {
			split := strings.SplitN(out.String(), "====== Report ======", 2)
			fmt.Fprint(&report, strings.TrimSpace(split[1]), "\n\n")
		} else {
			io.Copy(&out, &report)
		}

		utilities.Append(parameters.Output, report)
		io.Copy(os.Stdout, &out)
	}

	return err
}

func (c *Connection) Scp(source string, destination string) error {
	sftp, err := sftp.NewClient(c.Client)
	if err != nil {
		return err
	}
	defer sftp.Close()

	goEnum, err := os.Open(source)
	if err != nil {
		return err
	}
	defer goEnum.Close()

	destFile, err := sftp.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = destFile.ReadFrom(goEnum)

	return err
}

func (c *Connection) TemporaryFileCopy(source string, destination string, function func(string, string) error, params string) error {

	var (
		err error
		tmp error
	)

	err = c.Scp(source, destination)

	if err == nil {
		tmp = function(destination, params)
		if tmp != nil {
			err = tmp
		}
	}

	_, _, tmp = c.Shell(fmt.Sprintf("rm %v", destination))
	if tmp != nil {
		err = tmp
	}

	return err
}

func SSHExecuteGoEnum(ssh Connection) error {
	return ssh.TemporaryFileCopy(parameters.Binary, fmt.Sprintf("/tmp/%v", utilities.RandString(16)), ssh.executeGoEnum, parameters.RemoteParameters())
}

func SSHTmpfsExecuteGoEnum(ssh Connection) error {
	return ssh.Tmpfs(
		func(location string) error {
			return ssh.TemporaryFileCopy(parameters.Binary, fmt.Sprintf("%v/%v", location, utilities.RandString(16)), ssh.executeGoEnum, parameters.RemoteParameters())
		},
	)
}

func (c *Connection) executeGoEnum(destination string, params string) error {
	_, stderr, err := c.Shell(fmt.Sprintf("chmod +x %v", destination))
	if err != nil {
		return err
	}

	if stderr != "" {
		return errors.New(stderr)
	}

	err = c.ShellOutput(fmt.Sprintf("%v %v", destination, params))
	if err != nil {
		return err
	}
	return nil
}
