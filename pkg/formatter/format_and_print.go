package formatter

import (
	"sort"

	"github.com/DavydenkoGr/gitfame/pkg/author"
	"github.com/DavydenkoGr/gitfame/pkg/sorter"
)

func (f *Formatter) FormatAndPrint(authorDict author.AuthorDict) {
	authors := sorter.Authors{
		SortKey: f.OrderBy,
		Array: []*author.Author{},
	}

	for _, author := range authorDict {
		authors.Array = append(authors.Array, author)
	}

	sort.Sort(authors)

	switch f.Format {
	case "tabular":
		return
	case "csv":
		return
	case "json":
		return
	case "json-lines":
		return
	}
}
