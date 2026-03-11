package app

import (
	"fmt"
	"log/slog"

	"zen/pkg/app/config"
	cliflag "zen/pkg/app/flag"
	"zen/pkg/logs"
	"zen/pkg/signal"

	"github.com/spf13/cobra"
)

var appFlags cliflag.NamedFlagSets

type AppOption interface {
	Flags() cliflag.NamedFlagSets
}

var appOptions AppOption

func SetOption(opt AppOption) {
	appOptions = opt

	for _, fs := range opt.Flags().FlagSets {
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

	logger := logs.Init(logs.WithReplaceAttr(logAttrReplacer))
	slog.SetDefault(logger)

	if err := config.Init(cmd, appOptions); err != nil {
		return err
	}

	return nil
}
