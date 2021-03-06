package main

import (
	"strings"
	"testing"
)

func TestYAMLToJSON(t *testing.T) {
	indent = false
	revert = false
	expected := "{\"age\":30,\"name\":\"John\"}"
	actual := string(processFile("test/simple.yaml"))

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestYAMLToJSONIndent(t *testing.T) {
	indent = true
	revert = false
	expected := "{\n  \"age\": 30,\n  \"name\": \"John\"\n}"
	actual := string(processFile("test/simple.yaml"))

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestJSONToYAML(t *testing.T) {
	revert = true
	expected := "age: 30\nname: John\n"
	actual := string(processFile("test/simple.json"))

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
