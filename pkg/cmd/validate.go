package cmd

import (
	"fmt"
	"io"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/spf13/cobra"
	"github.com/vmware-tanzu/cartographer/pkg/apis/v1alpha1"
)

// ValidateCmd a struct for the validate command.
type ValidateCmd struct {
	File string
	Cmd  *cobra.Command
}

// NewValidateCmd creates a new validate command.
func NewValidateCmd() *cobra.Command {
	c := ValidateCmd{}

	cmd := &cobra.Command{
		Use:     "validate <file / or stdin>",
		Short:   "Validate a supply chain",
		Long:    "",
		Example: "validate <file / or stdin>",
		Aliases: []string{"check"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Run(cmd, args)
		},
		Args: cobra.MaximumNArgs(1),
	}

	return cmd
}

// Run runs the command.
func (c *ValidateCmd) Run(cmd *cobra.Command, args []string) error {
	// this does the trick
	var inputReader = cmd.InOrStdin()

	// the argument received looks like a file, we try to open it
	if len(args) > 0 && args[0] != "-" {
		file, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("failed open file: %v", err)
		}
		fmt.Println("reading from file")
		inputReader = file
	} else {
		fmt.Println("reading from stdin")
	}

	// we process the input reader, wherever to be his origin
	b, err := io.ReadAll(inputReader)
	if err != nil {
		return fmt.Errorf("failed process input: %v", err)
	}

	csc := v1alpha1.ClusterSupplyChain{}

	err = yaml.Unmarshal(b, &csc)
	if err != nil {
		return err
	}

	err = csc.ValidateCreate()
	if err != nil {
		return fmt.Errorf("unable to validate supply chain\n%s:\n%+v", string(b), err)
	}

	fmt.Println("OK")

	return nil
}
