package process

import "strings"

func (c *Context) ParseFile(filename string) {
	response, err := c.ExecuteCommand(
		"git", "blame", "--porcelain", c.Revision,
	)

	lines := strings.Split(response, "\n")
	for _, line := range lines {
		// TODO: parse file content
	}
}