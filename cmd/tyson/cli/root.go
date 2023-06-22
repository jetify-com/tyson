package cli

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
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
		log.Println(err)
		return 1
	}
	return 0
}

func Main() {
	code := Execute(context.Background(), os.Args[1:])
	os.Exit(code)
}
