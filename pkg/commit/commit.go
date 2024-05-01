package commit

type Commit struct {
	AuthorName string
	Lines      int
}

type CommitDict = map[string]*Commit
