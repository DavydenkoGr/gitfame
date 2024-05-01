package context

import (
	"os/exec"
	"strings"
)

// method executes console command and cut \n suffix from response
func (c *Context) ExecuteCommand(name string, args ...string) (string, error) {
	command := exec.Command(name, args...)
	command.Dir = c.CurrentDir
	result, err := command.Output()

	response := string(result)
	response = strings.TrimSuffix(response, "\n")

	return response, err
}
