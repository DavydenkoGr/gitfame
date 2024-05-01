package context

func (c *Context) ParseFileHeader(filename string) map[string]string {
	hash, _ := c.ExecuteCommand(
		"git", "log", "--format=%H", "-n 1", filename,
	)

	author, _ := c.ExecuteCommand(
		"git", "log", "--format=%an", "-n 1", filename,
	)

	return map[string]string{
		"hash": hash,
		"author": author,
	}
}
