package context

import "os/exec"

func (c *Context) ExecuteCommand(name string, args ...string) (string, error) {
	command := exec.Command(name, args...)
	command.Dir = c.CurrentDir
	result, err := command.Output()

	return string(result), err
}
