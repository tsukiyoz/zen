package flag

import (
	"bytes"
	goflag "flag"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/moby/term"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type NamedFlagSets struct {
	Order    []string
	FlagSets map[string]*pflag.FlagSet
}

func NewNamedFlagSets() *NamedFlagSets {
	return &NamedFlagSets{
		FlagSets: make(map[string]*pflag.FlagSet),
	}
}

func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = make(map[string]*pflag.FlagSet)
	}
	if _, ok := nfs.FlagSets[name]; !ok {
		flagSet := pflag.NewFlagSet(name, pflag.ContinueOnError)
		nfs.FlagSets[name] = flagSet
		nfs.Order = append(nfs.Order, name)
	}

	return nfs.FlagSets[name]
}

func (nfs *NamedFlagSets) AddFlagSet(fs *pflag.FlagSet) {
	if nfs.FlagSets == nil {
		nfs.FlagSets = make(map[string]*pflag.FlagSet)
	}
	if _, ok := nfs.FlagSets[fs.Name()]; !ok {
		nfs.FlagSets[fs.Name()] = fs
		nfs.Order = append(nfs.Order, fs.Name())
	}
}

func WordSepNormalizeFunc(fs *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

func InitFlags(fs *pflag.FlagSet) {
	fs.SetNormalizeFunc(WordSepNormalizeFunc)
	fs.AddGoFlagSet(goflag.CommandLine)
}

func PrintFlags(fs *pflag.FlagSet) {
	flagInfos := []any{}
	fs.VisitAll(func(flag *pflag.Flag) {
		flagInfos = append(flagInfos, flag.Name, flag.Value)
	})
	slog.Debug("FLAG all", flagInfos...)
}

const (
	usageFmt = "Usage:\n  %s\n"
)

func PrintSections(w io.Writer, fss NamedFlagSets, cols int) {
	for _, name := range fss.Order {
		fs := fss.FlagSets[name]
		if !fs.HasFlags() {
			continue
		}

		var buf bytes.Buffer
		fmt.Fprintf(&buf, "\n%s flags:\n\n%s", strings.ToUpper(name[:1])+name[1:], fs.FlagUsagesWrapped(cols))
		fmt.Fprint(w, buf.String())
	}
}

func SetUsageAndHelpFunc(cmd *cobra.Command, fss NamedFlagSets, cols int) {
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		PrintSections(cmd.OutOrStderr(), fss, cols)
		return nil
	})
	cmd.SetHelpFunc(func(c *cobra.Command, s []string) {
		fmt.Fprintf(cmd.OutOrStdout(), usageFmt, cmd.UseLine())
		PrintSections(cmd.OutOrStdout(), fss, cols)
	})
}

func TerminalSize(out io.Writer) (int, int, error) {
	outFd, isTerminal := term.GetFdInfo(out)
	if !isTerminal {
		return 0, 0, fmt.Errorf("given writer is no terminal")
	}
	winsize, err := term.GetWinsize(outFd)
	if err != nil {
		return 0, 0, err
	}
	return int(winsize.Width), int(winsize.Height), nil
}
