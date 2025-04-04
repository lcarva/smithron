package cmd

import (
	"fmt"
	"os"

	"github.com/lcarva/smithron/internal/forger"

	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

var forgeCommand = &cobra.Command{
	Use:   "forge",
	Short: "Forge a plan into an entity of execution",
	RunE: func(cmd *cobra.Command, args []string) error {
		planRaw, err := os.ReadFile(planPath)
		if err != nil {
			return fmt.Errorf("reading plan path at %q: %w", planPath, err)
		}

		var p forger.Plan
		if err := yaml.Unmarshal(planRaw, &p); err != nil {
			return fmt.Errorf("parsing plan: %w", err)
		}

		f, err := forger.GetForgerForTarget(target)
		if err != nil {
			return fmt.Errorf("getting forger: %w", err)
		}

		forged, err := f.Forge(cmd.Context(), p)
		if err != nil {
			return fmt.Errorf("forging plan: %w", err)
		}

		out := cmd.OutOrStdout()
		if _, err := out.Write(forged); err != nil {
			return fmt.Errorf("writin()g forged: %w", err)
		}

		return nil
	},
}

var (
	target   string
	planPath string
)

func init() {
	forgeCommand.Flags().StringVarP(&target, "target", "t", "", "Targeted CI provider.")
	if err := forgeCommand.MarkFlagRequired("target"); err != nil {
		panic(err)
	}

	forgeCommand.Flags().StringVarP(&planPath, "plan", "p", "", "Source plan to forge from")
	if err := forgeCommand.MarkFlagRequired("plan"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(forgeCommand)
}
