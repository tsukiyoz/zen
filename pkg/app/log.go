package app

import "log/slog"

const (
	timeFormat = "2006-01-02 15:04:05.000"
)

func logAttrReplacer(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "time" && a.Value.Kind() == slog.KindTime {
		v := a.Value.Time()
		a.Value = slog.StringValue(v.Format(timeFormat))
	}
	return a
}
