package runner

import (
	"time"

	cliflag "zen/pkg/app/flag"

	"github.com/spf13/pflag"
)

type Options struct {
	HTTP *HTTPOptions
	GRPC *GRPCOptions
}

func NewOptions() *Options {
	return &Options{
		HTTP: &HTTPOptions{
			Addr: ":8080",
		},
		GRPC: &GRPCOptions{
			Addr: ":9090",
		},
	}
}

func (s *Options) Flags() (fss cliflag.NamedFlagSets) {
	s.HTTP.AddFlags(fss.FlagSet("http"))
	s.GRPC.AddFlags(fss.FlagSet("grpc"))
	return
}

type HTTPOptions struct {
	Addr    string
	Timeout time.Duration
}

func (s *HTTPOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Addr, "http-addr", s.Addr, "The address to bind the HTTP server to.")
	fs.DurationVar(&s.Timeout, "http-timeout", s.Timeout, "The timeout for the HTTP server.")
}

type GRPCOptions struct {
	Addr string
}

func (s *GRPCOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Addr, "grpc-addr", s.Addr, "The address to bind the gRPC server to.")
}
