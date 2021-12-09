package jsonconvention

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type ConventionTest struct {
	FooBar string
	Foo    *string
	Bar    int `json:"overrideInt"`
	Person Person
	People []Person
}

type Person struct {
	Firstname string
	Lastname  string
	Birthday  time.Time
}

var people = []Person{{Firstname: "Foo", Lastname: "Bar"}, {Firstname: "Bar", Lastname: "Foo"}}

func TestConvention(t *testing.T) {
	model := ConventionTest{
		FooBar: "Hello world",
		Bar:    13,
		People: people,
	}
	jsonBytes, _ := Marshal(model, Convention)

	jsonStr := string(jsonBytes)

	fmt.Printf("Json: %s\n", jsonStr)

	if !strings.Contains(jsonStr, "FooBar_CONVENTION") ||
		!strings.Contains(jsonStr, "Foo_CONVENTION") ||
		!strings.Contains(jsonStr, "overrideInt") ||
		!strings.Contains(jsonStr, "\"Firstname_CONVENTION\":\"Foo\"") {
		t.Error("Marshal didn't return expected result")
	}
}

func TestUnmarshalConvention(t *testing.T) {
	json := []byte(`{"FooBar_CONVENTION":"Hello world","Foo_CONVENTION":"Hey hoo","overrideInt":13,"Person_CONVENTION":{"Firstname_CONVENTION":"","Lastname_CONVENTION":"","Birthday_CONVENTION":"0001-01-01T00:00:00Z"},"People_CONVENTION":[{"Firstname_CONVENTION":"Foo","Lastname_CONVENTION":"Bar","Birthday_CONVENTION":"0001-01-01T00:00:00Z"},{"Firstname_CONVENTION":"Bar","Lastname_CONVENTION":"Foo","Birthday_CONVENTION":"0001-01-01T00:00:00Z"}]}`)

	var model ConventionTest

	_ = Unmarshal(json, &model, Convention)

	fmt.Printf("Model: %#v\n", model)

	if model.FooBar != "Hello world" ||
		*model.Foo != "Hey hoo" ||
		model.Bar != 13 ||
		model.People[0].Firstname != "Foo" {
		t.Error("Unmarshal didn't return expected results")
	}
}

func TestDecodeConvention(t *testing.T) {
	input := strings.NewReader(`{"FooBar_CONVENTION":"Hello world","Foo_CONVENTION":"Hey hoo","overrideInt":13,"Person_CONVENTION":{"Firstname_CONVENTION":"First Person","Lastname_CONVENTION":"Last Person","Birthday_CONVENTION":"0001-01-01T00:00:00Z"},"People_CONVENTION":[{"Firstname_CONVENTION":"Foo","Lastname_CONVENTION":"Bar","Birthday_CONVENTION":"0001-01-01T00:00:00Z"},{"Firstname_CONVENTION":"Bar","Lastname_CONVENTION":"Foo","Birthday_CONVENTION":"0001-01-01T00:00:00Z"}]}`)

	var model ConventionTest

	decoder := NewDecoder(input)
	_ = decoder.Decode(&model, Convention)

	fmt.Printf("Model: %#v\n", model)

	if model.FooBar != "Hello world" ||
		*model.Foo != "Hey hoo" ||
		model.Bar != 13 ||
		model.People[0].Firstname != "Foo" {
		t.Error("Decode didn't return expected results")
	}
}

func Convention(name string) string {
	return name + "_CONVENTION"
}
