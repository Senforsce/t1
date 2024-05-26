# Template generation

To generate Go code from `*.t1` files, use the `t1` command line tool.

```
t1 generate
```

The `t1 generate` command will recurse into subdirectories and generate Go code for every `*.t1` file it finds.

The command will output a list of files that it processed, how long it took, and the total elapsed time.

```
main.t1 complete in 897.292µs
Generated code for 1 templates with 0 errors in 1.291292ms
```

## Advanced options

The `t1 generate` command has a `--help` option that prints advanced options.

These include the ability to generate code for a single file and to choose the number of parallel workers that `t1 generate` uses to create Go files.

By default `t1 generate` uses the number of CPUs that your machine has installed.

```
t1 generate --help
```

```
  -f string
        Optionally generates code for a single file, e.g. -f header.t1
  -help
        Print help and exit.
  -path string
        Generates code for all files in path. (default ".")
  -source-map-visualisations
        Set to true to generate HTML files to visualise the t1 code and its corresponding Go code.
  -w int
        Number of workers to run in parallel. (default runtime.NumCPU())
```
