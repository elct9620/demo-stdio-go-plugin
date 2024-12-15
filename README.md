Go Standard I/O Plugin System Example
===

This is an example of a plugin system built on top of standard I/O in Go.

## Usage

Clone the repository:

```bash
git clone https://github.com/elct9620/demo-stdio-go-plugin.git
```

Build the plugin:

```bash
make plugin-json
# make plugin-xml
```

> The binary will be generated in the `plugin-bin` directory.

Run the main program:

```bash
go run .
```

Change the plugin:

```bash
go run . -plugin xml
```
