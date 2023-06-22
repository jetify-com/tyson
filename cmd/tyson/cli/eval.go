package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"go.jetpack.io/tyson"
)

func EvalCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "eval <file.tson>",
		Args:          cobra.ExactArgs(1),
		Short:         "Evaluates a tson file and prints the result to stdout",
		RunE:          runCmd,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return command
}

func runCmd(cmd *cobra.Command, args []string) error {
	inputPath := args[0]
	bytes, err := tyson.Eval(inputPath)
	if err != nil {
		return err
	}

	err = printJSON(bytes)
	if err != nil {
		return err
	}
	return nil
}

func printJSON(bytes []byte) error {
	if !isTerminal() {
		color.NoColor = true
	}
	output, err := prettyjson.Format(bytes)
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}

func isTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}
