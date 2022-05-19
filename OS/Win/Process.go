//go:build windows

package Win

import (
	"errors"
	"fmt"
	"os/exec"
)

// KillProcess kills a process in Windows
func KillProcess(progName string) (err error) {
	kill := exec.Command("taskkill", "/t", "/f", "/im", progName)
	err = kill.Run()
	var exerr *exec.ExitError // https://github.com/golang/go/issues/35874
	if errors.As(err, &exerr) || err == nil {
		return nil
	} else {
		return fmt.Errorf("could not kill process %s", progName)
	}
}
