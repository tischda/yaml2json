package main

import (
	"encoding/json"
	"log"

	"os"

	"bytes"

	"github.com/ghodss/yaml"
)

func convert(filename string, cfg *Config) (b []byte) {
	in, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	if cfg.revert {
		b, err = yaml.JSONToYAML(in)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		b, err = yaml.YAMLToJSON(in)
		if err != nil {
			log.Fatalln(err)
		}
		if cfg.indent {
			// wish ghodss/yaml had a pretty print option
			var prettyJSON bytes.Buffer
			if err = json.Indent(&prettyJSON, b, "", "  "); err != nil {
				log.Fatalln(err)
			}
			b = prettyJSON.Bytes()
		}
	}
	return
}
