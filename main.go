package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"io/ioutil"

	"os"

	"bytes"

	"github.com/ghodss/yaml"
)

const name string = "yaml2json"

var version string

var showVersion bool
var revert bool
var indent bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&indent, "indent", true, "indent JSON")
	flag.BoolVar(&revert, "revert", false, "transform JSON back to YAML")
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [option] filename\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n OPTIONS:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if flag.Arg(0) == "version" || showVersion {
		fmt.Printf("%s version %s\n", name, version)
		return
	}
	if flag.NArg() != 1 {
		flag.Usage()
	}
	out := processFile(flag.Arg(0))
	fmt.Print(string(out))
}

func processFile(name string) (b []byte) {
	in, err := ioutil.ReadFile(name)
	checkFatal(err)
	if revert {
		b, err = yaml.JSONToYAML(in)
		checkFatal(err)
	} else {
		b, err = yaml.YAMLToJSON(in)
		checkFatal(err)
		if indent {
			// wish ghodss/yaml had a pretty print option
			var prettyJSON bytes.Buffer
			err = json.Indent(&prettyJSON, b, "", "  ")
			checkFatal(err)
			b = prettyJSON.Bytes()
		}
	}
	return
}

func checkFatal(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
