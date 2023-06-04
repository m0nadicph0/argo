package cmd

import (
	"github.com/m0nadicph0/argo/internal/engine"
	"github.com/m0nadicph0/argo/internal/parser"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "argo",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		eng := &engine.Engine{Parser: parser.NewParser()}
		err := eng.Run(args)

		if err != nil {
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().SetInterspersed(false)
}
