package context

import (
	"github.com/DavydenkoGr/gitfame/pkg/author"
	"github.com/DavydenkoGr/gitfame/pkg/commit"
)

// main method which collects info about developers activities
func (c *Context) ParseRepository() author.AuthorDict {
	filenames := c.ParseTree()

	authorDict := make(author.AuthorDict)
	authorToCommit := make(map[string]map[string]struct{})
	authorToFilename := make(map[string]map[string]struct{})

	for _, filename := range filenames {
		header := c.ParseFileHeader(filename)
		contentCommitDict:= c.ParseFileContent(filename)
		
		if _, ok := contentCommitDict[header["hash"]]; !ok {
			contentCommitDict[header["hash"]] = &commit.Commit{
				AuthorName: header["author"],
				Lines: 0,
			}
		}

		for hash, commit := range contentCommitDict {

			if _, ok := authorDict[commit.AuthorName]; !ok {
				authorDict[commit.AuthorName] = &author.Author{
					Name: commit.AuthorName,
					Lines: 0,
					Commits: 0,
					Files: 0,
				}

				authorToCommit[commit.AuthorName] = map[string]struct{}{}
				authorToFilename[commit.AuthorName] = map[string]struct{}{}
			}

			authorDict[commit.AuthorName].Lines += commit.Lines

			authorToCommit[commit.AuthorName][hash] = struct{}{}
			authorToFilename[commit.AuthorName][filename] = struct{}{}
		}
	}

	for author, commitDict := range authorToCommit {
		authorDict[author].Commits += len(commitDict)
	}
	
	for author, filenameDict := range authorToFilename {
		authorDict[author].Files += len(filenameDict)
	}

	return authorDict
}
