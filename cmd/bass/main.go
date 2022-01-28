package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"

	"github.com/spf13/cobra"
	"github.com/vito/bass/pkg/bass"
	"github.com/vito/bass/pkg/ioctx"
	"github.com/vito/bass/pkg/runtimes"
	"github.com/vito/bass/pkg/zapctx"
)

var rootCmd = &cobra.Command{
	Use:           "bass [scriptfile args]",
	Long:          "run a bass script, or start a repl (if no args are given)",
	Version:       bass.Version,
	Example:       `bass ci/bass`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          root,
}

var profPort int
var profFilePath string

func init() {
	rootCmd.Flags().IntVar(&profPort, "profile", 0, "port number to bind for Go HTTP profiling")
	rootCmd.Flags().StringVar(&profFilePath, "cpu-profile", "", "take a CPU profile and save it to this path")
}

func main() {
	logger := bass.Logger()
	ctx := zapctx.ToContext(context.Background(), logger)

	trace := &bass.Trace{}
	ctx = bass.WithTrace(ctx, trace)

	ctx = ioctx.StderrToContext(ctx, os.Stderr)

	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

var DefaultConfig = bass.Config{
	Runtimes: []bass.RuntimeConfig{
		{
			Platform: bass.LinuxPlatform,
			Runtime:  runtimes.BuildkitName,
		},
	},
}

func root(cmd *cobra.Command, argv []string) error {
	ctx := cmd.Context()

	if profPort != 0 {
		zapctx.FromContext(ctx).Sugar().Debugf("serving pprof on :%d", profPort)

		l, err := net.Listen("tcp", fmt.Sprintf(":%d", profPort))
		if err != nil {
			bass.WriteError(ctx, err)
			return err
		}

		go http.Serve(l, nil)
	}

	if profFilePath != "" {
		profFile, err := os.Create(profFilePath)
		if err != nil {
			bass.WriteError(ctx, err)
			return err
		}

		defer profFile.Close()

		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}

	ctx, err := initCtx(ctx)
	if err != nil {
		// NB: root() is responsible for printing its own errors so main() doesn't
		// do it redundantly with the console UI. initialization steps are broken
		// out into initCtx() just to make it harder to "miss a spot" and fail
		// silently.
		bass.WriteError(ctx, err)
		return err
	}

	if runExport {
		return export(ctx)
	}

	if runPrune {
		return prune(ctx)
	}

	if runLSP {
		return langServer(ctx)
	}

	if len(argv) == 0 {
		return repl(ctx)
	}

	return run(ctx, argv[0], argv[1:]...)
}

func initCtx(ctx context.Context) (context.Context, error) {
	config, err := bass.LoadConfig(DefaultConfig)
	if err != nil {
		return nil, err
	}

	pool, err := runtimes.NewPool(config)
	if err != nil {
		return nil, err
	}

	return bass.WithRuntimePool(ctx, pool), nil
}
