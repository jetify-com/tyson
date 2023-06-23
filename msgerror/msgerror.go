package msgerror

// TODO: there's an opportunity to bundle this code, together with:
// - the code from esbuild that creates and formats error messages
// - the error sanitization code we have in devbox
// to opensource a library around friendly error messages
// Something similar to rust's https://docs.rs/color-eyre/latest/color_eyre/
import (
	"errors"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/mattn/go-isatty"
)

type Error struct {
	err      error
	messages []esbuild.Message
}

func ErrFromMessages(toplevel string, messages []esbuild.Message) error {
	return &Error{
		err:      errors.New(toplevel),
		messages: messages,
	}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Messages() []string {
	isTerminal := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	formatted := api.FormatMessages(e.messages, api.FormatMessagesOptions{
		Kind:          api.ErrorMessage,
		Color:         isTerminal,
		TerminalWidth: 100,
	})
	return formatted
}
