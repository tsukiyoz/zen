package app

import (
	"fmt"

	"zen/pkg/app/config"
	cliflag "zen/pkg/app/flag"
	"zen/pkg/logs"
	"zen/pkg/signal"

	"github.com/spf13/cobra"
)

var appFlags cliflag.NamedFlagSets

// RegisterFlags register custom flags to appFlags, it should run before CompleteCommand.
func RegisterFlags(fss cliflag.NamedFlagSets) {
	for _, fs := range fss.FlagSets {
		appFlags.AddFlagSet(fs)
	}
}

// CompleteCommand add basic flags to appFlags, and normalize cliflags.
func CompleteCommand(cmd *cobra.Command) {
	globalFlags := appFlags.FlagSet("global")
	logs.AddFlags(globalFlags)
	config.AddFlags(globalFlags)
	globalFlags.BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))

	for _, fs := range appFlags.FlagSets {
		cmd.Flags().AddFlagSet(fs)
	}

	cols, _, _ := cliflag.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, appFlags, cols)
}

func PreRun(cmd *cobra.Command, args []string) {
	_ = PreRunE(cmd, args)
}

func PreRunE(cmd *cobra.Command, args []string) error {
	ctx := signal.SetupSignalContext()
	cmd.SetContext(ctx)

	logs.Init()

	if err := config.Init(cmd); err != nil {
		return err
	}

	return nil
}
