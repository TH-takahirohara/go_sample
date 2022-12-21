package main

import "testing"

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
