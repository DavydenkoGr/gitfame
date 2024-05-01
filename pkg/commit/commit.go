package commit

// single file commit information
type Commit struct {
	AuthorName string
	Lines      int
}

type CommitDict = map[string]*Commit
