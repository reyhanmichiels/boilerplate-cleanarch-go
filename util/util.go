package util

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CreateFile(name string) error {
	_, err := os.Stat(name)

	if os.IsExist(err) {

		return errors.New("file already exist")

	}

	if strings.Contains(name, "/") {

		path := name[:strings.LastIndex(name, "/")]

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {

			return err

		}

	}

	_, err = os.Create(name)
	if err != nil {

		return err

	}

	return nil
}

func ExecuteCommand(command string) (string, error) {
	var output []byte
	var err error

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", command).Output()
	} else {
		output, err = exec.Command("bash", "-c", command).Output()
	}

	return string(output), err
}
