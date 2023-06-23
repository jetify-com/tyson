package cli

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.jetpack.io/tyson/msgerror"
)

func RootCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "tyson",
		Short: "TypeScript as a configuration language",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	command.AddCommand(EvalCmd())

	return command
}

func Execute(ctx context.Context, args []string) int {
	cmd := RootCmd()
	cmd.SetArgs(args)
	err := cmd.ExecuteContext(ctx)
	if err != nil {
		var msgError *msgerror.Error
		if errors.As(err, &msgError) {
			for _, msg := range msgError.Messages() {
				fmt.Fprintln(os.Stderr, msg)
			}
		} else {
			fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err)
		}
		return 1
	}
	return 0
}

func Main() {
	code := Execute(context.Background(), os.Args[1:])
	os.Exit(code)
}
