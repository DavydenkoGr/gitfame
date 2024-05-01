package cmd

import (
	// "fmt"
	// "os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
	"github.com/DavydenkoGr/gitfame/pkg/context"
	"github.com/DavydenkoGr/gitfame/pkg/filter"
	"github.com/DavydenkoGr/gitfame/pkg/formatter"
)

var (
	// Used for flags.
	appFilter = &filter.Filter{}
	appFormatter = &formatter.Formatter{}
	appContext = &context.Context{
		Filter: appFilter,
		Formatter: appFormatter,
	}

	rootCmd = &cobra.Command{
		Use:   "gitfame",
		Short: "gitfame",
		Long: "gitfame",
		Run: func(cmd *cobra.Command, args []string) {
			authorDict := appContext.ParseRepository()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&appContext.CurrentDir, "repository", ".", "repository")
	rootCmd.PersistentFlags().StringVar(&appContext.Revision, "revision", "HEAD", "revision")

	if *rootCmd.PersistentFlags().Bool("use-committer", false, "use-committer") {
		appContext.AuthorType = context.CommitterT
	} else {
		appContext.AuthorType = context.AuthorT
	}

	rootCmd.PersistentFlags().StringSliceVar(&appFilter.Extensions, "extensions", []string{}, "extensions")
	rootCmd.PersistentFlags().StringSliceVar(&appFilter.Languages, "languages", []string{}, "languages")
	rootCmd.PersistentFlags().StringSliceVar(&appFilter.Exclude, "exclude", []string{}, "exclude")
	rootCmd.PersistentFlags().StringSliceVar(&appFilter.RestrictTo, "restrict-to", []string{}, "restrict-to")

	rootCmd.PersistentFlags().StringVar(&appFormatter.Format, "format", "tabular", "format")
	rootCmd.PersistentFlags().StringVar(&appFormatter.OrderBy, "order-by", "lines", "order-by")

	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".cobra")
// 	}

// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
