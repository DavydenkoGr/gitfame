package process

func (c *Context) ParseFileHeader(filename string) map[string]string {
	hash, err := c.ExecuteCommand(
		"git", "log", "--format=%H", "-n 1", filename,
	)

	author, err := c.ExecuteCommand(
		"git", "log", "--format=%an", "-n 1", filename,
	)

	return map[string]string{
		"hash": hash,
		"author": author,
	}
}
