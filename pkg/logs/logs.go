package logs

import (
	"log/slog"
	"os"
	"sync"

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

var (
	initOnce sync.Once
	logger   *slog.Logger
)

func Init(opts ...Option) *slog.Logger {
	initOnce.Do(func() {
		opt := &options{
			level:  logLevel,
			format: format,
		}

		for _, o := range opts {
			o(opt)
		}

		var level slog.Level
		err := level.UnmarshalText([]byte(opt.level))
		if err != nil {
			panic("failed to init log")
		}
		sopt := &slog.HandlerOptions{
			Level:     slog.Level(level),
			AddSource: opt.addSource,
		}

		if opt.replaceAttr != nil {
			sopt.ReplaceAttr = opt.replaceAttr
		}

		var handler slog.Handler

		if opt.handler != nil {
			handler = opt.handler
		} else {
			switch opt.format {
			case "json":
				handler = slog.NewJSONHandler(os.Stdout, sopt)
			case "text":
				handler = slog.NewTextHandler(os.Stdout, sopt)
			}
		}
		if handler == nil {
			handler = slog.NewTextHandler(os.Stdout, sopt)
		}

		logger = slog.New(handler)
	})

	return logger
}

func AddFlags(fs *pflag.FlagSet) {
	packageFlags.VisitAll(func(f *pflag.Flag) {
		if fs.Lookup(f.Name) == nil {
			fs.AddFlag(f)
		}
	})
}
