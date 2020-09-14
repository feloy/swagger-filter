package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "swagger-filter",
	Short: "Filter a swagger specification by endpoint",
	Long:  `Filter a swagger specification by endpoint`,
	RunE:  filterCmd,
	Args:  cobra.MinimumNArgs(2),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.swagger-filter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().StringSliceP("endpoint", "", nil, "Endpoints to filter, by exact path")
	rootCmd.Flags().StringSliceP("endpoint-prefix", "", nil, "Endpoints to filter, by prefix")
	rootCmd.Flags().StringSliceP("endpoint-regexp", "", nil, "Endpoints to filter, by regexp")
}
