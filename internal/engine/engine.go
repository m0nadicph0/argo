package engine

import (
	"github.com/m0nadicph0/argo/internal/parser"
	"github.com/m0nadicph0/argo/internal/util"
	"os"
	"os/exec"
)

type Engine struct {
	Parser               parser.Parser
	MaxArgsPerInvocation int
}

func (e Engine) Run(args []string) error {
	// Execute the command with the arguments
	inArgs, _ := e.Parser.GetTokens()
	invocations := util.GroupBySize(inArgs, e.MaxArgsPerInvocation)
	for _, invocationArgs := range invocations {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Args = append(cmd.Args, invocationArgs...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return err
		}

	}
	return nil
}
