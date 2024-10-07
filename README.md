![t1](https://github.com/senforsce/tndr/raw/main/t1.png)

## A HTML templating language for Go that has great developer tooling.

![t1](ide-demo.gif)

## Documentation

See user documentation at https://senforsce.com/t1/guide

<p align="center">
<a href="https://pkg.go.dev/github.com/senforsce/tndr"><img src="https://pkg.go.dev/badge/github.com/senforsce/tndr.svg" alt="Go Reference" /></a>
<a href="https://xcfile.dev"><img src="https://xcfile.dev/badge.svg" alt="xc compatible" /></a>
<a href="https://raw.githack.com/wiki/senforsce/t1/coverage.html"><img src="https://github.com/senforsce/tndr/wiki/coverage.svg" alt="Go Coverage" /></a>
<a href="https://goreportcard.com/report/github.com/senforsce/tndr"><img src="https://goreportcard.com/badge/github.com/senforsce/tndr" alt="Go Report Card" /></a<
</p>

## Tasks

### build

Build a local version.

```sh
go run ./get-version > .version
cd cmd/t1
go build
```

### install-snapshot

Build and install current version.

```sh
# Remove t1 from the non-standard ~/bin/t1 path
# that this command previously used.
rm -f ~/bin/t1
# Clear LSP logs.
rm -f cmd/t1/lspcmd/*.txt
# Update version.
go run ./get-version > .version
# Install to $GOPATH/bin or $HOME/go/bin
cd cmd/t1 && go install
```

### build-snapshot

Use goreleaser to build the command line binary using goreleaser.

```sh
goreleaser build --snapshot --clean
```

### generate

Run t1 generate using local version.

```sh
go run ./cmd/t1 generate -include-version=false
```

### test

Run Go tests.

```sh
go run ./get-version > .version
go run ./cmd/t1 generate -include-version=false
go test ./...
```

### test-cover

Run Go tests.

```sh
# Create test profile directories.
mkdir -p coverage/fmt
mkdir -p coverage/generate
mkdir -p coverage/version
mkdir -p coverage/unit
# Build the test binary.
go build -cover -o ./coverage/t1-cover ./cmd/t1
# Run the covered generate command.
GOCOVERDIR=coverage/fmt ./coverage/t1-cover fmt .
GOCOVERDIR=coverage/generate ./coverage/t1-cover generate -include-version=false
GOCOVERDIR=coverage/version ./coverage/t1-cover version
# Run the unit tests.
go test -cover ./... -args -test.gocoverdir="$PWD/coverage/unit"
# Display the combined percentage.
go tool covdata percent -i=./coverage/fmt,./coverage/generate,./coverage/version,./coverage/unit
# Generate a text coverage profile for tooling to use.
go tool covdata textfmt -i=./coverage/fmt,./coverage/generate,./coverage/version,./coverage/unit -o coverage.out
# Print total
go tool cover -func coverage.out | grep total
```

### benchmark

Run benchmarks.

```sh
go run ./cmd/t1 generate -include-version=false && go test ./... -bench=. -benchmem
```

### fmt

Format all Go and t1 code.

```sh
gofmt -s -w .
go run ./cmd/t1 fmt .
```

### lint

```sh
golangci-lint run --verbose
```

### release

Create production build with goreleaser.

```sh
if [ "${GITHUB_TOKEN}" == "" ]; then echo "No github token, run:"; echo "export GITHUB_TOKEN=`pass github.com/goreleaser_access_token`"; exit 1; fi
./push-tag.sh
goreleaser --clean
```

### docs-run

Run the development server.

Directory: docs

```sh
npm run start
```

### docs-build

Build production docs site.

Directory: docs

```sh
npm run build
```
