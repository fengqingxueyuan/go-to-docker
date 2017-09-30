package cmd

import (
	"flag"
	// "fmt"
	// "log"
	// "os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	// "google.golang.org/grpc"

	// "k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-docker/metering/pkg/server"
)

func RootCommandFor(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Metering application for cAdvisor/Heapster/Prometheus/Docker-stats metrics",
		Long: `
        metering
        
        This is a client of cAdvisor and any metrics system, and serving to export
        metrics data into destination as time series,
        It is inspired by jaeger and prometheus.
        `,
	}
	root.AddCommand(createExporterCommand())
	root.AddCommand(createCollectorCommand())
	// root.AddCommand(createClientCommand())

	return root
}

func createExporterCommand() *cobra.Command {
	var loglevel string

	command := &cobra.Command{
		Use:   "export",
		Short: "Start exporter serving with streaming and a gRPC-Gateway for interactive stuff",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()
			server.RunExporter()
		},
	}

	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

func createCollectorCommand() *cobra.Command {
	var loglevel string

	command := &cobra.Command{
		Use:   "collect",
		Short: "Start collector serving with streaming and a gRPC-Gateway for interactive stuff",
		Run: func(cmd *cobra.Command, args []string) {
			// pflag.Parse()
			flag.Set("v", loglevel)
			flag.Parse()
			server.RunCollector()
		},
	}

	command.Flags().StringVar(&loglevel, "loglevel", "2", "for glog")
	// command.Flags().AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return command
}

func createClientCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "client",
	}
	return command
}
