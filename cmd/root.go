package cmd

import (
	"github.com/m0nadicph0/argo/internal/engine"
	"github.com/m0nadicph0/argo/internal/parser"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "argo",
	Short: "Constructs argument list and executes utility commands",
	Run: func(cmd *cobra.Command, args []string) {
		maxArgsPerInvocation, _ := cmd.Flags().GetInt("num-max-arg")
		eng := &engine.Engine{
			Parser:               parser.NewParser(),
			MaxArgsPerInvocation: maxArgsPerInvocation,
		}
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
	rootCmd.Flags().IntP("num-max-arg", "n", 5000, "Maximum number of arguments taken from standard input for each invocation")
}
