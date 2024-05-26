# Command-line tools

`t1` provides a command line interface. Most users will only need to run the `t1 generate` command to generate Go code from `*.t1` files.

```
usage: t1 <command> [parameters]
To see help text, you can run:
  t1 generate --help
  t1 fmt --help
  t1 lsp --help
  t1 migrate --help
  t1 version
examples:
  t1 generate
```

## Generating Go code from t1 files

The `t1 generate` command generates Go code from `*.t1` files in the current directory tree.

The command provides additional options:

```
usage: t1 generate [<args>...]

Generates Go code from t1 files.

Args:
  -path <path>
    Generates code for all files in path. (default .)
  -f <file>
    Optionally generates code for a single file, e.g. -f header.t1
  -sourceMapVisualisations
    Set to true to generate HTML files to visualise the t1 code and its corresponding Go code.
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
```

For example, to generate code for a single file:

```
t1 generate -f header.t1
```

## Formatting t1 files

The `t1 fmt` command formats template files. You can use this command in different ways:

1. Format all template files in the current directory and subdirectories:

```
t1 fmt .
```

2. Format input from stdin and output to stdout:

```
t1 fmt
```

## Language Server for IDE integration

`t1 lsp` provides a Language Server Protocol (LSP) implementation to support IDE integrations.

This command isn't intended to be used directly by users, but is used by IDE integrations such as the VSCode extension and by Neovim support.

A number of additional options are provided to enable runtime logging and profiling tools.

```
  -goplsLog string
        The file to log gopls output, or leave empty to disable logging.
  -goplsRPCTrace
        Set gopls to log input and output messages.
  -help
        Print help and exit.
  -http string
        Enable http debug server by setting a listen address (e.g. localhost:7474)
  -log string
        The file to log t1 LSP output to, or leave empty to disable logging.
  -pprof
        Enable pprof web server (default address is localhost:9999)
```
