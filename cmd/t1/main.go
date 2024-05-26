package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"

	"github.com/fatih/color"
	"github.com/senforsce/t1"
	"github.com/senforsce/t1/cmd/t1/fmtcmd"
	"github.com/senforsce/t1/cmd/t1/generatecmd"
	"github.com/senforsce/t1/cmd/t1/lspcmd"
)

func main() {
	code := run(os.Stdout, os.Args)
	if code != 0 {
		os.Exit(code)
	}
}

const usageText = `usage: t1 <command> [<args>...]

Tndr - build Semantic Web With OWL & HTMX UIs with Go

See docs at https://senforsce.com/t1/guide and https://senforsce.com/tndr/guide (coming soon)
Based off https://t1.guide

commands:
  generate   Generates Go code from t1 or t1 files
  fmt        Formats t1 or t1 files
  lsp        Starts a language server for t1 or t1 files
  version    Prints the version
`

func run(w io.Writer, args []string) (code int) {
	if len(args) < 2 {
		fmt.Fprint(w, usageText)
		return 0
	}
	switch args[1] {
	case "generate":
		return generateCmd(w, args[2:])
	case "fmt":
		return fmtCmd(w, args[2:])
	case "lsp":
		return lspCmd(w, args[2:])
	case "version":
		fmt.Fprintln(w, t1.Version())
		return 0
	case "--version":
		fmt.Fprintln(w, t1.Version())
		return 0
	}
	fmt.Fprint(w, usageText)
	return 0
}

const generateUsageText = `usage: t1 generate [<args>...]

Generates Go code from t1 or t1 files.

Args:
  -path <path>
    Generates code for all files in path. (default .)
  -f <file>
    Optionally generates code for a single file, e.g. -f header.t1
  -sourceMapVisualisations
    Set to true to generate HTML files to visualise the t1 or t1 code and its corresponding Go code.
  -include-version
    Set to false to skip inclusion of the t1 version in the generated code. (default true)
  -include-timestamp
    Set to true to include the current time in the generated code.
  -watch
    Set to true to watch the path for changes and regenerate code.
  -cmd <cmd>
    Set the command to run after generating code.
  -proxy
    Set the URL to proxy after generating code and executing the command.
  -proxyport
    The port the proxy will listen on. (default 7331)
  -w
    Number of workers to use when generating code. (default runtime.NumCPUs)
  -pprof
    Port to run the pprof server on.
  -keep-orphaned-files
    Keeps orphaned generated t1 files. (default false)
  -v
    Set log verbosity level to "debug". (default "info")
  -log-level
    Set log verbosity level. (default "info", options: "debug", "info", "warn", "error")
  -help
    Print help and exit.

Examples:

  Generate code for all files in the current directory and subdirectories:

    t1 generate

  Generate code for a single file:

    t1 generate -f header.t1

  Watch the current directory and subdirectories for changes and regenerate code:

    t1 generate -watch
`

func generateCmd(w io.Writer, args []string) (code int) {
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)
	cmd.SetOutput(w)
	fileNameFlag := cmd.String("f", "", "")
	pathFlag := cmd.String("path", ".", "")
	sourceMapVisualisationsFlag := cmd.Bool("source-map-visualisations", false, "")
	includeVersionFlag := cmd.Bool("include-version", true, "")
	includeTimestampFlag := cmd.Bool("include-timestamp", false, "")
	watchFlag := cmd.Bool("watch", false, "")
	openBrowserFlag := cmd.Bool("open-browser", true, "")
	cmdFlag := cmd.String("cmd", "", "")
	proxyFlag := cmd.String("proxy", "", "")
	proxyPortFlag := cmd.Int("proxyport", 7331, "")
	workerCountFlag := cmd.Int("w", runtime.NumCPU(), "")
	pprofPortFlag := cmd.Int("pprof", 0, "")
	keepOrphanedFilesFlag := cmd.Bool("keep-orphaned-files", false, "")
	verboseFlag := cmd.Bool("v", false, "")
	logLevelFlag := cmd.String("log-level", "info", "")
	helpFlag := cmd.Bool("help", false, "")
	err := cmd.Parse(args)
	if err != nil || *helpFlag {
		fmt.Fprint(w, generateUsageText)
		return
	}

	logLevel := *logLevelFlag
	if *verboseFlag {
		logLevel = "debug"
	}

	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		fmt.Fprintln(w, "Stopping...")
		cancel()
	}()
	err = generatecmd.Run(ctx, w, generatecmd.Arguments{
		FileName:                        *fileNameFlag,
		Path:                            *pathFlag,
		Watch:                           *watchFlag,
		OpenBrowser:                     *openBrowserFlag,
		Command:                         *cmdFlag,
		Proxy:                           *proxyFlag,
		ProxyPort:                       *proxyPortFlag,
		WorkerCount:                     *workerCountFlag,
		GenerateSourceMapVisualisations: *sourceMapVisualisationsFlag,
		IncludeVersion:                  *includeVersionFlag,
		IncludeTimestamp:                *includeTimestampFlag,
		LogLevel:                        logLevel,
		PPROFPort:                       *pprofPortFlag,
		KeepOrphanedFiles:               *keepOrphanedFilesFlag,
	})
	if err != nil {
		color.New(color.FgRed).Fprint(w, "(✗) ")
		fmt.Fprintln(w, err.Error())
		return 1
	}
	return 0
}

const fmtUsageText = `usage: t1 fmt [<args> ...]

Format all files in directory:

  t1 fmt .

Format stdin to stdout:

  t1 fmt < header.t1

Format file or directory to stdout:

  t1 fmt -stdout FILE

Args:
  -stdout
    Prints to stdout instead of in-place format
  -v
    Set log verbosity level to "debug". (default "info")
  -log-level
    Set log verbosity level. (default "info", options: "debug", "info", "warn", "error")
  -w
    Number of workers to use when formatting code. (default runtime.NumCPUs).
  -help
    Print help and exit.
`

func fmtCmd(w io.Writer, args []string) (code int) {
	cmd := flag.NewFlagSet("fmt", flag.ExitOnError)
	cmd.SetOutput(w)
	cmd.Usage = func() {
		fmt.Fprint(w, fmtUsageText)
	}
	helpFlag := cmd.Bool("help", false, "")
	workerCountFlag := cmd.Int("w", runtime.NumCPU(), "")
	verboseFlag := cmd.Bool("v", false, "")
	logLevelFlag := cmd.String("log-level", "info", "")
	stdout := cmd.Bool("stdout", false, "")

	err := cmd.Parse(args)
	if err != nil || *helpFlag {
		cmd.Usage()
		return
	}

	logLevel := *logLevelFlag
	if *verboseFlag {
		logLevel = "debug"
	}

	err = fmtcmd.Run(w, fmtcmd.Arguments{
		ToStdout:    *stdout,
		Files:       cmd.Args(),
		LogLevel:    logLevel,
		WorkerCount: *workerCountFlag,
	})
	if err != nil {
		return 1
	}
	return 0
}

const lspUsageText = `usage: t1 lsp [<args> ...]

Starts a language server for t1.

Args:
  -log string
    The file to log t1 LSP output to, or leave empty to disable logging.
  -goplsLog string
    The file to log gopls output, or leave empty to disable logging.
  -goplsRPCTrace
    Set gopls to log input and output messages.
  -help
    Print help and exit.
  -pprof
    Enable pprof web server (default address is localhost:9999)
  -http string
    Enable http debug server by setting a listen address (e.g. localhost:7474)
`

func lspCmd(w io.Writer, args []string) (code int) {
	cmd := flag.NewFlagSet("lsp", flag.ExitOnError)
	cmd.SetOutput(w)
	log := cmd.String("log", "", "")
	goplsLog := cmd.String("goplsLog", "", "")
	goplsRPCTrace := cmd.Bool("goplsRPCTrace", false, "")
	helpFlag := cmd.Bool("help", false, "")
	pprofFlag := cmd.Bool("pprof", false, "")
	httpDebugFlag := cmd.String("http", "", "")
	err := cmd.Parse(args)
	if err != nil || *helpFlag {
		fmt.Fprint(w, lspUsageText)
		return
	}
	err = lspcmd.Run(w, lspcmd.Arguments{
		Log:           *log,
		GoplsLog:      *goplsLog,
		GoplsRPCTrace: *goplsRPCTrace,
		PPROF:         *pprofFlag,
		HTTPDebug:     *httpDebugFlag,
	})
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return 1
	}
	return 0
}
