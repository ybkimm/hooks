package cmd

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "hooks",
	Short: "@ybkimm's claude-code hooks collection.",
}

func Get() *cobra.Command {
	return root
}

func AddCommand(cmds ...*cobra.Command) {
	root.AddCommand(cmds...)
}
