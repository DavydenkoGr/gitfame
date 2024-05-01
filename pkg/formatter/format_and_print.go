package formatter

import (
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

	authors.Sort()

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
