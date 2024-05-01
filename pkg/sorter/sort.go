package sorter

import (
	"github.com/DavydenkoGr/gitfame/pkg/author"
)

type Authors struct {
	SortKey string
	Array []*author.Author
}

func (a Authors) Len() int {
	return len(a.Array)
}

func (a Authors) Swap(i, j int) {
	a.Array[i], a.Array[j] = a.Array[j], a.Array[i]
}

func (a Authors) Less(i, j int) bool {
	first, second := a.Array[i], a.Array[j]

	switch a.SortKey {
	case "lines":
		if first.Lines < second.Lines {
			return true
		}
	case "commits":
		if first.Commits < second.Commits {
			return true
		}
	case "files":
		if first.Files < second.Files {
			return true
		}
	}

	if first.Lines < second.Lines {
		return true
	}

	if first.Commits < second.Commits {
		return true
	}

	if first.Files < second.Files {
		return true
	}

	if first.Name < second.Name {
		return true
	}

	return false
}
