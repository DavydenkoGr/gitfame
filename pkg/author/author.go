package author

type Author struct {
	Name    string `json:"name"`
	Lines   int `json:"lines"`
	Commits int `json:"commits"`
	Files   int `json:"files"`
}

type AuthorDict = map[string]*Author

type AuthorType int
const (
	AuthorT AuthorType = iota
	CommitterT
)
