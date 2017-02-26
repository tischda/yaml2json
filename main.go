package main

import (
	"flag"
	"fmt"
	"log"

	"io/ioutil"

	"os"

	"github.com/ghodss/yaml"
)

const name string = "yaml2json"

var version string

var showVersion bool
var revert bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
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
	}
	return
}

func checkFatal(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
