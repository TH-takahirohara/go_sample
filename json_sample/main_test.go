package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestUnmarshalFruits(t *testing.T) {
	testStr := `
	[{
		"name":"apple", "color":"red"
	},{
		"name":"banana", "color":"yellow"
	}]
	`
	wants := []Fruit{{Name: "apple", Color: "red"}, {Name: "banana", Color: "yellow"}}

	fruits, err := UnmarshalFruits(testStr)
	if err != nil {
		t.Fatalf("failed to unmarshal input string: %v", err)
	}
	for i, f := range fruits {
		if f != wants[i] {
			t.Fatal("return value is not correct")
		}
	}
}

func TestMarshalFruits(t *testing.T) {
	fruits := []Fruit{
		{Name: "apple", Color: "red"},
		{Name: "banana", Color: "yellow"},
	}
	wants := `[{"name":"apple","color":"red"},{"name":"banana","color":"yellow"}]`

	get, err := MarshalFruits(fruits)
	if err != nil {
		t.Fatalf("failed to marshal input fruits objects: %v", err)
	}
	if wants != get {
		t.Fatalf("return value is not correct")
	}
}

func TestUnmarshalFruitsFile(t *testing.T) {
	path := "testdata/test_fruits.json"
	wants := []Fruit{{Name: "apple", Color: "red"}, {Name: "banana", Color: "yellow"}}

	getFruits, err := UnmarshalFruitsFile(path)
	if err != nil {
		t.Fatalf("failed to unmarshal input text: %v", err)
	}

	for i, f := range getFruits {
		if f != wants[i] {
			t.Fatalf("return value is not correct")
		}
	}
}

func TestEncodeFruits(t *testing.T) {
	fruits := []Fruit{
		{Name: "apple", Color: "red"},
		{Name: "banana", Color: "yellow"},
	}
	wants := `[{"name":"apple","color":"red"},{"name":"banana","color":"yellow"}]`
	b := new(bytes.Buffer)

	err := EncodeFruits(fruits, b)
	if err != nil {
		t.Fatalf("cannot encode: %v", err)
	}

	if strings.ReplaceAll(b.String(), "\n", "") != wants {
		t.Fatalf("wanted string is %v, but got string is %v", wants, b.String())
	}
}

func TestDecodeFruits(t *testing.T) {
	fruits := new([]Fruit)
	wants := []Fruit{{Name: "apple", Color: "red"}, {Name: "banana", Color: "yellow"}}
	s := `[{"name":"apple","color":"red"},{"name":"banana","color":"yellow"}]`

	err := DecodeFruits(fruits, s)
	if err != nil {
		t.Fatalf("cannot decode: %v", err)
	}

	for i, v := range *fruits {
		if v != wants[i] {
			t.Fatalf("got %v, but want %v", v, wants[i])
		}
	}
}
