package author

// structure which contains info about author or committer activities
type Author struct {
	Name    string `json:"name"`
	Lines   int `json:"lines"`
	Commits int `json:"commits"`
	Files   int `json:"files"`
}

type AuthorDict = map[string]*Author

// type of author
type AuthorType int
const (
	AuthorT AuthorType = iota
	CommitterT
)
