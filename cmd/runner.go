/*
Copyright © 2026 tsukiyo <tsukiyo6@163.com>
*/
package cmd

import (
	"log/slog"

	"zen/cmd/options/runner"
	"zen/internal"
	"zen/pkg/app"

	"github.com/spf13/cobra"
)

// runnerCmd represents the runner command
var runnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "runner",
	Long:  `this is runner`,

	PreRunE: app.PreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Info("runner called")

		return internal.Run(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(runnerCmd)

	s := runner.NewOptions()

	app.RegisterFlags(s.Flags())
	app.CompleteCommand(runnerCmd)
}
