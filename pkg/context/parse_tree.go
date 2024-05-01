package context

import "strings"

// method parses git repo tree 
func (c *Context) ParseTree() []string {
	response, _ := c.ExecuteCommand(
		"git", "ls-tree", "--name-only", "-r", c.Revision,
	)

	result := make([]string, 0)
	for _, filename := range strings.Split(response, "\n") {

		if c.Filter.IsFiltered(filename) {
			result = append(result, filename)
		}
	}

	return result
}
