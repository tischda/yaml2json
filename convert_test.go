package main

import (
	"strings"
	"testing"
)

func TestYAMLToJSON(t *testing.T) {
	cfg := &Config{}

	expected := "{\"age\":30,\"name\":\"John\"}"
	actual := string(convert("test/simple.yaml", cfg))

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestYAMLToJSONIndent(t *testing.T) {
	cfg := &Config{
		indent: true,
	}
	expected := "{\n  \"age\": 30,\n  \"name\": \"John\"\n}"
	actual := string(convert("test/simple.yaml", cfg))

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

func TestJSONToYAML(t *testing.T) {
	cfg := &Config{
		revert: true,
	}

	expected := "age: 30\nname: John\n"
	actual := string(convert("test/simple.json", cfg))

	if strings.Compare(actual, expected) != 0 {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
