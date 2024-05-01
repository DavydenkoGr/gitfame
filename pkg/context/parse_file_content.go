package context

import "strings"

// method parses file content and count every author lines for all commits
func (c *Context) ParseFileContent(filename string) CommitDict {
	response, err := c.ExecuteCommand(
		"git", "blame", "--porcelain", c.Revision,
	)

	commitDict := make(CommitDict)
	// empty file
	if len(response) == 0 {
		return commitDict
	}

	lines := strings.Split(response, "\n")
	lineIndex := 0

	// parse commits info
	for line := lines[lineIndex]; lineIndex < len(lines); lineIndex++ {
		if c.AuthorType == AuthorT && strings.HasPrefix(line, "author ") {
			hash := strings.Split(lines[lineIndex - 1], " ")[0]

			commitDict[hash] = &Commit{
				AuthorName: line[7:],
				Lines: 1,
			}

			continue
		}

		if c.AuthorType == CommitterT && strings.HasPrefix(line, "committer ") {
			hash := strings.Split(lines[lineIndex - 5], " ")[0]

			commitDict[hash] = &Commit{
				AuthorName: line[10:],
				Lines: 1,
			}

			continue
		}

		// count lines
		hash := strings.Split(line, " ")[0]
		if _, ok := commitDict[hash]; ok {
			commitDict[hash].Lines++
		}
	}
	
	return commitDict
}
