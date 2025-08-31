[![Build Status](https://github.com/tischda/yaml2json/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/yaml2json/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/yaml2json/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/yaml2json/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tischda/yaml2json)](https://goreportcard.com/report/github.com/tischda/yaml2json)

# yaml2json

Converts Packer configuration files written in YAML to the required JSON format.

The main point is to write comments in the configuration file and HCL is not ready yet.

See also:
   * https://github.com/mitchellh/packer/issues/1768
   * https://github.com/mitchellh/packer/pull/4461


### Install

There is a dependency on github.com/ghodss/yaml.

~~~
go install github.com/tischda/yaml2json@latest
~~~

### Usage

~~~
Usage: yaml2json.exe [option] filename

 OPTIONS:
  -indent
        indent JSON (default true)
  -revert
        transform JSON back to YAML
  -version
        print version and exit
~~~

Example:

~~~
# cat test\simple.yaml
name: John
age: 30

# yaml2json.exe test\simple.yaml
{
  "age": 30,
  "name": "John"
}
~~~
