package gitfame.process

import os/exec

struct Filter {
	// TODO: implement me
}


type Context struct {
	Revision string
	CurrentDir string
	Filter Filter
	// TODO: add new fields
}

type Author struct {
	Name string
	Lines int
	Commits int
	Files int
}

func (f *Filter) IsFiltered(string) {
	// TODO: implement me
}

func (c *Context) ExecuteCommand(command string, args ...string) {
	command := exec.Command(command, args...)
	exec.Dir(c.CurrentDir)
	return command.Output()
}
