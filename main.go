package main

import (
	"flag"
	"fmt"
	"log"

	"os"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	indent  bool
	revert  bool
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.indent, "i", true, "")
	flag.BoolVar(&cfg.indent, "indent", true, "indent JSON")
	flag.BoolVar(&cfg.revert, "r", false, "")
	flag.BoolVar(&cfg.revert, "revert", false, "transform JSON back to YAML")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] filename

Converts configuration files written in YAML to JSON format.

OPTIONS:
  -i, --indent
        indent JSON output (default: true)
  -r, --revert
        transform JSON back to YAML
  -?, --help
        display this help message
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` test\simple.yaml
  {
    "age": 30,
    "name": "John"
  }`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	filename := flag.Arg(0)
	out := convert(filename, cfg)
	fmt.Print(string(out))
}
