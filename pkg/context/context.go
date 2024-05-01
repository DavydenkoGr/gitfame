package context

import (
	"github.com/DavydenkoGr/gitfame/pkg/filter"
	"github.com/DavydenkoGr/gitfame/pkg/formatter"
)

type AuthorType int
const (
	AuthorT AuthorType = iota
	CommitterT
)

type Context struct {
	AuthorType AuthorType
	Revision   string
	CurrentDir string
	Filter     *filter.Filter
	Formatter  *formatter.Formatter
}
