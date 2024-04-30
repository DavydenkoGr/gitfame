package gitfame.process


import strings

func (c *Context) ParseTree() []string {
	result, err := c.ExecuteCommand(
		"git", "ls-tree", "--name-only", "-r", c.Revision
	)

	result := make([]string)
	for filename := range strings.Split(result, "\n") {
		if c.Filter(filename) {
			result = append(result, filename)
		}
	}

	return result
}
