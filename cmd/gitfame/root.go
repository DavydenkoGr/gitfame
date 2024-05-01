package cmd

import (
	"github.com/DavydenkoGr/gitfame/pkg/author"
	"github.com/DavydenkoGr/gitfame/pkg/context"
	"github.com/DavydenkoGr/gitfame/pkg/filter"
	"github.com/DavydenkoGr/gitfame/pkg/formatter"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	appFilter = filter.Filter{}
	appFormatter = formatter.Formatter{}
	appContext = context.Context{
		Filter: &appFilter,
		Formatter: &appFormatter,
	}

	rootCmd = &cobra.Command{
		Use:   "gitfame",
		Short: "gitfame",
		Long: "gitfame",
		Run: func(cmd *cobra.Command, args []string) {
			appContext.RunApp()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

// gitfame initialization
func init() {
	// appContext fields init
	rootCmd.PersistentFlags().StringVar(&appContext.CurrentDir, "repository", ".", "repository")
	rootCmd.PersistentFlags().StringVar(&appContext.Revision, "revision", "HEAD", "revision")

	if *rootCmd.PersistentFlags().Bool("use-committer", false, "use-committer") {
		appContext.AuthorType = author.CommitterT
	} else {
		appContext.AuthorType = author.AuthorT
	}

	// appFilter fields init
	extensions := []string{}
	languages := []string{}
	rootCmd.PersistentFlags().StringSliceVar(&extensions, "extensions", []string{}, "extensions")
	rootCmd.PersistentFlags().StringSliceVar(&languages, "languages", []string{}, "languages")

	for _, extension := range extensions {
		appFilter.ExtensionsSet[extension] = struct{}{}
	}

	appFilter.InitializeLanguages(languages)

	rootCmd.PersistentFlags().StringSliceVar(&appFilter.Exclude, "exclude", []string{}, "exclude")
	rootCmd.PersistentFlags().StringSliceVar(&appFilter.RestrictTo, "restrict-to", []string{}, "restrict-to")

	// appFormatter fields init
	rootCmd.PersistentFlags().StringVar(&appFormatter.Format, "format", "tabular", "format")
	rootCmd.PersistentFlags().StringVar(&appFormatter.OrderBy, "order-by", "lines", "order-by")
}
