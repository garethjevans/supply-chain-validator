package cmd

import (
	"errors"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vmware-tanzu/cartographer/pkg/apis/v1alpha1"
)

// ValidateCmd a struct for the validate command.
type ValidateCmd struct {
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
			err := c.Run()
			if err != nil {
				logrus.Fatalf("%s", err)
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

	err = csc.ValidateCreate()
	if err != nil {
		return fmt.Errorf("unable to validate supply chain\n%s:\n%+v", string(b), err)
	}

	fmt.Println("OK")

	return nil
}
