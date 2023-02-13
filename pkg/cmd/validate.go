package cmd

import (
	"github.com/fatih/color"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	colorError = color.New(color.FgRed).SprintFunc()
)

// ValidateCmd a struct for the validate command.
type ValidateCmd struct {
	Cmd  *cobra.Command
	Args []string
}

// NewValidateCmd creates a new validate command.
func NewValidateCmd() *cobra.Command {
	c := ValidateCmd{}

	cmd := &cobra.Command{
		Use:     "validate -f <path>",
		Short:   "Validate a supply chain",
		Long:    "",
		Example: "validate -f <path>",
		Aliases: []string{"check"},
		Run: func(cmd *cobra.Command, args []string) {
			c.Cmd = cmd
			c.Args = args
			err := c.Run()
			if err != nil {
				logrus.Fatalf("unable to run command: %s", err)
			}
		},
		Args: cobra.NoArgs,
	}

	return cmd
}

// Run runs the command.
func (c *ValidateCmd) Run() error {
	return nil
}
