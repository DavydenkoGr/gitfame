package author

type Author struct {
	Name    string
	Lines   int
	Commits int
	Files   int
}

type AuthorDict = map[string]*Author

type AuthorType int
const (
	AuthorT AuthorType = iota
	CommitterT
)
