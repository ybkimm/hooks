package posttoolgo

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
	"github.com/ybkimm/hooks/cmd"
	"github.com/ybkimm/hooks/internal/hooks"
	"github.com/ybkimm/hooks/internal/iox"
)

var MutableTools = []string{
	"Edit",
	"MultiEdit",
	"Write",
}

func init() {
	cmd.AddCommand(&cobra.Command{
		Use:   "postedit-go",
		Short: "Post-edit hook for go language",
		Run: func(cmd *cobra.Command, args []string) {
			ok, err := run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if !ok {
				os.Exit(2)
			}
		},
	})
}

func run() (bool, error) {
	input, err := hooks.ReadInput[hooks.FileInput, struct{}](os.Stdin)
	if err != nil {
		return false, fmt.Errorf("failed to decode input: %v", err)
	}

	var (
		isPostToolEvent = input.Event == "PostToolUse"
		isMutableTool   = slices.Contains(MutableTools, input.Tool)
		isGoFile        = strings.HasSuffix(input.ToolInput.Path, ".go")
		isFileExists    = iox.Exists(input.ToolInput.Path)
	)

	if !isPostToolEvent || !isMutableTool || !isGoFile || !isFileExists {
		return true, nil
	}

	var ok = true

	// 1. Format
	_, err = script.Exec(fmt.Sprintf(
		"gopls format -w %s",
		input.ToolInput.Path,
	)).
		WithStderr(os.Stderr).
		Stdout()
	if err != nil {
		return false, fmt.Errorf("gopls fmt failed: %v", err)
	}

	// 2. Sort imports
	_, err = script.Exec(fmt.Sprintf(
		"gopls imports -w %s",
		input.ToolInput.Path,
	)).
		WithStderr(os.Stderr).
		Stdout()
	if err != nil {
		return false, fmt.Errorf("gopls imports failed: %v", err)
	}

	// 3. Check error, warnings, info
	_, err = script.Exec(fmt.Sprintf("gopls check -severity=hint %s", input.ToolInput.Path)).
		Filter(func(r io.Reader, w io.Writer) error {
			n, err := io.Copy(w, r)
			if err != nil {
				return err
			}
			if n > 0 {
				ok = false
			}
			return nil
		}).
		WithStdout(os.Stderr).
		WithStderr(os.Stderr).
		Stdout()
	if err != nil {
		return false, fmt.Errorf("gopls check failed: %v", err)
	}

	// Done!

	return ok, nil
}
