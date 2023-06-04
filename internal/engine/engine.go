package engine

import (
	"github.com/m0nadicph0/argo/internal/parser"
	"os"
	"os/exec"
)

type Engine struct {
	Parser parser.Parser
}

func (e Engine) Run(args []string) error {
	// Execute the command with the arguments
	inArgs, _ := e.Parser.GetTokens()
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Args = append(cmd.Args, inArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
