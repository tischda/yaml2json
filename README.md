[![Build Status](https://github.com/tischda/yaml2json/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/yaml2json/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/yaml2json/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/yaml2json/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/yaml2json/badge.svg)](https://coveralls.io/r/tischda/yaml2json)
[![Linter Status](https://github.com/tischda/yaml2json/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/yaml2json/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/yaml2json)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/yaml2json.svg)](https://github.com/tischda/yaml2json/releases/latest)


# yaml2json

Converts configuration files written in YAML to the required JSON format.

The main point is to write comments in Packer configuration files and HCL is not ready yet.

See also:
   * https://github.com/mitchellh/packer/issues/1768
   * https://github.com/mitchellh/packer/pull/4461


## Install

There is a dependency on `github.com/ghodss/yaml`.

~~~
go install github.com/tischda/yaml2json@latest
~~~

## Usage

~~~
Usage: yaml2json [OPTIONS] filename

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
~~~

## Examples

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
