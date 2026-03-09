package logs

import (
	"log/slog"
	"os"

	"github.com/spf13/pflag"
)

var (
	packageFlags = pflag.NewFlagSet("log", pflag.ContinueOnError)
	logLevel     string
	format       string
)

func init() {
	packageFlags.StringVar(&logLevel, "log-level", "info", "log level")
	packageFlags.StringVar(&format, "log-format", "json", "log output format(text,json)")
}

const (
	timeFormat = "2006-01-02 15:04:05.000"
)

func Init() {
	var level slog.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		panic("failed to init log")
	}
	opts := &slog.HandlerOptions{
		Level:     slog.Level(level),
		AddSource: false,
	}

	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}

func AddFlags(fs *pflag.FlagSet) {
	packageFlags.VisitAll(func(f *pflag.Flag) {
		if fs.Lookup(f.Name) == nil {
			fs.AddFlag(f)
		}
	})
}
