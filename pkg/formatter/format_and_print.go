package formatter

import (
	"encoding/csv"
	"fmt"
	"github.com/DavydenkoGr/gitfame/pkg/author"
	"github.com/DavydenkoGr/gitfame/pkg/sorter"
	"os"
	"strconv"
	"text/tabwriter"
)

func (f *Formatter) FormatAndPrint(authorDict author.AuthorDict) {
	authors := sorter.Authors{
		SortKey: f.OrderBy,
		Array:   []*author.Author{},
	}

	for _, author := range authorDict {
		authors.Array = append(authors.Array, author)
	}

	authors.Sort()

	switch f.Format {
	case "tabular":
		printTabular(authors)
	case "csv":
		printCSV(authors)
	case "json":
		return
	case "json-lines":
		return
	}
}

func printTabular(authors sorter.Authors) {
	tabwriter := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	fmt.Fprintln(tabwriter, "Name\tLines\tCommits\tFiles")

	for _, author := range authors.Array {
		fmt.Fprintf(
			tabwriter,
			"%s\t%d\t%d\t%d\n",
			author.Name,
			author.Lines,
			author.Commits,
			author.Files,
		)
	}

	tabwriter.Flush()
}

func printCSV(authors sorter.Authors) {
	records := [][]string{
		{"Name", "Lines", "Commits", "Files"},
	}

	for _, author := range authors.Array {
		records = append(
			records,
			[]string{
				author.Name,
				strconv.Itoa(author.Lines),
				strconv.Itoa(author.Commits),
				strconv.Itoa(author.Files),
			},
		)
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)
}
