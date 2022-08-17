package remote

import (
	"errors"
	"fmt"
	"os"

	"github.com/goEnum/goEnum/utilities"
)

func (c *Connection) Tmpfs(function func(string) error) error {
	var location string
	for {
		location = fmt.Sprintf("/tmp/%v", utilities.RandString(16))

		_, err := os.Stat(location)
		if errors.Is(err, os.ErrNotExist) {
			break
		}
	}
	var (
		err error
		tmp error
	)
	_, _, err = c.Shell("mkdir " + location)
	if err == nil {
		_, _, tmp = c.Shell("sudo mount -t tmpfs -o size=10m swap " + location)
		if tmp != nil {
			err = tmp
		}

		if err == nil {
			tmp = function(location)
			if tmp != nil {
				err = tmp
			}
		}

		_, _, tmp = c.Shell("sudo umount " + location)
		if tmp != nil {
			err = tmp
		}
	}
	_, _, tmp = c.Shell("rmdir " + location)
	if tmp != nil {
		err = tmp
	}

	return err
}
