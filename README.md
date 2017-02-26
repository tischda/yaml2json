# yaml2json [![Build status](https://ci.appveyor.com/api/projects/status/l3a401e7voipe8sa/branch/master?svg=true)](https://ci.appveyor.com/project/tischda/yaml2json/branch/master)

Small utility written in [Go](https://www.golang.org) to convert Packer configuration files
written in YAML to the required JSON format.

The main reason is that I need to write comments in this configuration file and HCL is not ready yet.

See also:
   * https://github.com/mitchellh/packer/issues/1768
   * https://github.com/mitchellh/packer/pull/4461


### Install

There is a dependency on github.com/ghodss/yaml.

~~~
go get github.com/kardianos/govendor

git clone https://github.com/tischda/yaml2json

cd yaml2json

govendor sync
make dist
~~~

### Usage

~~~
Usage: yaml2json.exe [option] filename

 OPTIONS:
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
{"age":30,"name":"John"}
~~~
