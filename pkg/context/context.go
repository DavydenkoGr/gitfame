package context

import (
	"github.com/DavydenkoGr/gitfame/pkg/author"
	"github.com/DavydenkoGr/gitfame/pkg/filter"
	"github.com/DavydenkoGr/gitfame/pkg/formatter"
)

type Context struct {
	AuthorType author.AuthorType
	Revision   string
	CurrentDir string
	Filter     *filter.Filter
	Formatter  *formatter.Formatter
}
