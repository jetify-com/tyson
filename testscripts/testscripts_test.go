package testscripts

import (
	"context"
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"go.jetpack.io/tyson/cmd/tyson/cli"
)

func TestScripts(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                 "testdata",
		RequireExplicitExec: true,
		RequireUniqueNames:  true,
		TestWork:            false,
	})
}

func TestMain(m *testing.M) {
	commands := map[string]func() int{
		"tyson": func() int {
			return cli.Execute(context.Background(), os.Args[1:])
		},
	}
	os.Exit(testscript.RunMain(m, commands))
}
