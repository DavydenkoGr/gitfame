package context

func (c *Context) ParseRepository() AuthorDict {
	filenames := c.ParseTree()

	authorToCommit := make(map[string]map[string]struct{})
	authorDict := make(AuthorDict)

	for _, filename := range filenames {
		header := c.ParseFileHeader(filename)
		contentCommitDict:= c.ParseFileContent(filename)
		
		if _, ok:= contentCommitDict[header["hash"]]; !ok {
			contentCommitDict["hash"] = &Commit{
				AuthorName: header["author"],
				Lines: 0,
			}
		}

		for hash, commit := range contentCommitDict {
			authorDict[commit.AuthorName].Lines += commit.Lines
			authorDict[commit.AuthorName].Files++

			authorToCommit[commit.AuthorName][hash] = struct{}{}
		}
	}

	for author, commits := range authorToCommit {
		authorDict[author].Commits += len(commits)
	}

	return authorDict
}
