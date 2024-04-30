package process

import "strings"

func (c *Context) ParseFile(filename string) CommitDict {
	response, err := c.ExecuteCommand(
		"git", "blame", "--porcelain", c.Revision,
	)

	commitDict := make(CommitDict)
	if len(response) == 0 {
		return commitDict
	}

	lines := strings.Split(response, "\n")
	lineIndex := 0

	// parse commits info
	for line := lines[lineIndex]; lineIndex < len(lines); lineIndex++ {
		if len(line) == 0 {
			break
		}

		if c.AuthorType == AuthorT && strings.HasPrefix(line, "author ") {
			hash := strings.Split(lines[lineIndex - 1], " ")[0]

			commitDict[hash] = &Commit{
				AuthorName: line[7:],
			}
		}

		if c.AuthorType == CommitterT && strings.HasPrefix(line, "committer ") {
			hash := strings.Split(lines[lineIndex - 5], " ")[0]

			commitDict[hash] = &Commit{
				AuthorName: line[10:],
			}
		}
	}

	for line := lines[lineIndex]; lineIndex < len(lines); lineIndex += 2 {
		hash := strings.Split(line, " ")[0]
		commitDict[hash].Lines++
	}
	
	return commitDict
}
