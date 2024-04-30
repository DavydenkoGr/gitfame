package process

func (c *Context) ParseFileHeader(filename string) CommitDict {
	hash, err := c.ExecuteCommand(
		"git", "log", "--format=%H", "-n 1", filename,
	)

	author, err := c.ExecuteCommand(
		"git", "log", "--format=%an", "-n 1", filename,
	)

	cd := make(CommitDict)
	cd[hash] = &Commit{
		AuthorName: author,
		Lines: 0,
	}

	return cd
}
