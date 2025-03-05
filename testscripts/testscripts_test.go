package testscripts

import (
	"context"
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"go.jetify.com/tyson/cmd/tyson/cli"
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
	commands := map[string]func(){
		"tyson": func() {
			os.Exit(cli.Execute(context.Background(), os.Args[1:]))
		},
	}
	testscript.Main(m, commands)
}
