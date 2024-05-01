package context

type Author struct {
	Name    string
	Lines   int
	Commits int
	Files   int
}

type AuthorDict = map[string]*Author
