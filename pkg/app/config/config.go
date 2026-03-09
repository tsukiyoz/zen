package config

import (
	"errors"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	packageFlags    = pflag.NewFlagSet("config", pflag.ContinueOnError)
	cfgFile         string
	cfgFileFlagName = "config"
)

func init() {
	packageFlags.StringVarP(&cfgFile, "config", "c", "", "config file path")
}

func AddFlags(fs *pflag.FlagSet) {
	packageFlags.VisitAll(func(f *pflag.Flag) {
		if fs.Lookup(f.Name) == nil {
			fs.AddFlag(f)
		}
	})
}

func Init(cmd *cobra.Command, opt any) error {
	slog.Debug("Loading env configurations...", "component", cmd.Name())
	viper.SetEnvPrefix(cmd.Name())
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(".")
		viper.AddConfigPath(path.Join(home, cmd.Name()))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var cfgFileNotFoundErr viper.ConfigFileNotFoundError
		if !errors.As(err, &cfgFileNotFoundErr) {
			return err
		}
	}

	if usedConfigFile := viper.ConfigFileUsed(); usedConfigFile != "" {
		slog.Debug("Configuration initialized", "config file", viper.ConfigFileUsed())
	} else {
		slog.Warn("Config file not found")
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	if opt != nil {
		if err := viper.Unmarshal(opt); err != nil {
			slog.Error("unmarshal configs failed", "err", err)
		}
	}

	return nil
}
