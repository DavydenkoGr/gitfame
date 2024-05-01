package process

import "os/exec"

type Filter struct {
	// TODO: implement me
}

type AuthorType int
const (
	AuthorT AuthorType = iota
	CommitterT
)

type Context struct {
	AuthorType AuthorType
	Revision   string
	CurrentDir string
	Filter     *Filter
	// TODO: add new fields
}

type Author struct {
	Name    string
	Lines   int
	Commits int
	Files   int
}

type Commit struct {
	AuthorName string
	Lines      int
}

type CommitDict = map[string]*Commit
type AuthorDict = map[string]*Author

func (f *Filter) IsFiltered(string) bool {
	// TODO: implement me

	return true
}

func (c *Context) ExecuteCommand(name string, args ...string) (string, error) {
	command := exec.Command(name, args...)
	command.Dir = c.CurrentDir
	result, err := command.Output()

	return string(result), err
}
