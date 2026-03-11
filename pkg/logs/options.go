package logs

import "log/slog"

type Option func(*options)

type options struct {
	level       string
	format      string
	handler     slog.Handler
	addSource   bool
	replaceAttr func(groups []string, a slog.Attr) slog.Attr
}

func WithLevel(level string) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithFormat(format string) Option {
	return func(o *options) {
		o.format = format
	}
}

func WithHandler(h slog.Handler) Option {
	return func(o *options) {
		o.handler = h
	}
}

func WithAddSource(add bool) Option {
	return func(o *options) {
		o.addSource = add
	}
}

func WithReplaceAttr(f func(groups []string, a slog.Attr) slog.Attr) Option {
	return func(o *options) {
		o.replaceAttr = f
	}
}
