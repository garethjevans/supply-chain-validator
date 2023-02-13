package cmd

import (
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vmware-tanzu/cartographer/pkg/apis/v1alpha1"
	"os"
)

// ValidateCmd a struct for the validate command.
type ValidateCmd struct {
	//Cmd  *cobra.Command
	//Args []string
	File string
}

// NewValidateCmd creates a new validate command.
func NewValidateCmd() *cobra.Command {
	c := ValidateCmd{}

	cmd := &cobra.Command{
		Use:     "validate -f <file>",
		Short:   "Validate a supply chain",
		Long:    "",
		Example: "validate -f <file>",
		Aliases: []string{"check"},
		Run: func(cmd *cobra.Command, args []string) {
			//c.Cmd = cmd
			//c.Args = args
			err := c.Run()
			if err != nil {
				logrus.Fatalf("unable to run command: %s", err)
			}
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().StringVarP(&c.File, "file", "f", "", "Path to a file containing a cluster supply chain resource")

	return cmd
}

// Run runs the command.
func (c *ValidateCmd) Run() error {
	if _, err := os.Stat(c.File); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("unable to find file %s", c.File)
	}

	csc := v1alpha1.ClusterSupplyChain{}

	b, err := os.ReadFile(c.File)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &csc)
	if err != nil {
		return err
	}

	fmt.Println("OK")

	return nil
}
